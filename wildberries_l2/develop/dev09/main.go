package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Generat17/dev09/getfile"
	"github.com/Generat17/dev09/wq"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"path"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	maxWorkers      = 20
	defaultMaxDepth = 3
)

type depthURL struct {
	url   *url.URL
	depth int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("url must be provided")
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	outputDir := "."
	if len(os.Args) >= 3 {
		outputDir = os.Args[2]
	}

	maxDepth := defaultMaxDepth
	if len(os.Args) >= 4 {
		maxDepth, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := wget(ctx, url, maxDepth, outputDir); err != nil {
		log.Fatal(err)
	}
}

func wget(ctx context.Context, URL *url.URL, maxDepth int, outputDir string) error {
	wq := wq.New(maxWorkers)

	linkCh := make(chan depthURL)
	ulinkCh := make(chan depthURL)

	go uniqueLinkFilter(ulinkCh, linkCh, URL)

	go func() {
		for link := range ulinkCh {
			wq.AddTask(wgetPageTask(link, maxDepth, outputDir, linkCh))
		}
	}()

	wq.AddTask(wgetPageTask(depthURL{URL, 0}, maxDepth, outputDir, linkCh))
	wq.RunAndWait(ctx)
	close(linkCh)
	return nil
}

func wgetPageTask(link depthURL, maxDepth int, outputDir string, linkCh chan<- depthURL) wq.Task {
	return func(ctx context.Context) {
		err := wgetPage(ctx, link, maxDepth, outputDir, linkCh)
		if err != nil {
			log.Println(err)
		}
	}
}

func wgetPage(ctx context.Context, URL depthURL, maxDepth int, outputDir string, linkCh chan<- depthURL) error {
	body, fileType, err := getfile.GetFile(URL.url)
	if err != nil {
		return err
	}
	defer body.Close()

	file, err := prepareFile(URL.url, outputDir)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("dowloading:", URL.url.String())

	if fileType == getfile.HTML && URL.depth < maxDepth {
		err = linkTransformer(ctx, file, body, URL, linkCh)
	} else {
		// do not handle links if any
		_, err = io.Copy(file, body)
	}

	// fmt.Println("done dowloading:", URL.String())
	return err
}

func prepareFile(url *url.URL, outputDir string) (*os.File, error) {
	filepath := path.Join(outputDir, url.Path)
	dirpath := path.Dir(filepath)

	err := os.MkdirAll(dirpath, 0775)
	if err != nil {
		// check if dir exists as regular file
		stat, serr := os.Stat(dirpath)
		if serr == nil && stat.Mode().IsRegular() {
			_ = os.Rename(dirpath, dirpath+".tmp")
			_ = os.MkdirAll(dirpath, 0775)
			_ = os.Rename(dirpath+".tmp", dirpath+"/index.html")
		} else {
			return nil, err
		}
	}

	file, err := os.Create(filepath)
	if err != nil {
		// check if filepath already is a directory
		stat, serr := os.Stat(filepath)
		if serr == nil && stat.Mode().IsDir() {
			filepath += "/index.html"
			return os.Create(filepath)
		}
	}
	return file, err
}

func linkTransformer(ctx context.Context, w io.Writer, r io.Reader, baseURL depthURL, urlCh chan<- depthURL) error {
	urlRe, err := urlMatcher(baseURL.url)
	if err != nil {
		return err
	}

	sc := bufio.NewReader(r)
	for ctx.Err() == nil {
		line, err := sc.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		linkGroups := urlRe.FindAllStringSubmatch(line, -1)

		for _, links := range linkGroups {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			url, err := parseURL(links[1], baseURL.url)
			if err != nil {
				return err
			}
			urlCh <- depthURL{url, baseURL.depth + 1}
			line = strings.ReplaceAll(line, links[1], url.Path)
		}

		_, werr := io.WriteString(w, line+"\n")
		if werr != nil {
			return err
		}

		if err == io.EOF {
			return nil
		}
	}

	return ctx.Err()
}

func uniqueLinkFilter(out chan<- depthURL, in <-chan depthURL, baseURL *url.URL) {
	defer close(out)

	seen := map[string]struct{}{
		indexise(baseURL.Path): {},
	}

	for url := range in {
		key := indexise(url.url.Path)
		if _, ok := seen[key]; !ok {
			out <- url
			seen[key] = struct{}{}
		}
	}
}

func parseURL(urlStr string, hostURL *url.URL) (*url.URL, error) {
	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = hostURL.Scheme + "://" + hostURL.Host + urlStr
	}
	return url.Parse(urlStr)
}

func urlMatcher(url *url.URL) (*regexp.Regexp, error) {
	host := strings.ReplaceAll(url.Hostname(), ".", `\.`)
	urlReStr := fmt.Sprintf(`(?:href|src)="((https://[a-zA-Z.]*%s)?(/[^"]+?)?)"`, host)
	return regexp.Compile(urlReStr)
}

func indexise(path string) string {
	if !strings.HasSuffix(path, "/index.html") {
		path += "/index.html"
	}
	return path
}
