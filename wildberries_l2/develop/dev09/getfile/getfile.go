package getfile

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// FileType - тип скачиваемого файла (для пост обработки)
type FileType int

const (
	// HTML - гипертекстовый документ
	HTML FileType = iota
	// Other - другой тип
	Other
)

// GetFile - получение файла по сети
func GetFile(URL *url.URL) (io.ReadCloser, FileType, error) {
	resp, err := http.Get(URL.String())
	if err != nil {
		return nil, 0, err
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return resp.Body, HTML, nil
	}
	return resp.Body, Other, nil
}
