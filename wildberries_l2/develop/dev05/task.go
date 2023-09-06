package main

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	errUnknownFlag     = errors.New("unknown flag")
	errValueExpected   = errors.New("value expected")
	errPatternExpected = errors.New("pattern expected")
)

type args struct {
	after      int
	before     int
	count      int
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func main() {
	args, err := parseArgs(os.Args[1:])
	if err == nil {
		err = runGrep(os.Stdout, os.Stdin, args)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func runGrep(w io.Writer, r io.Reader, args args) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	windowCh := make(chan []string)
	go windowedReader(ctx, windowCh, r, args.before, args.after)

	lineNumber := 0
	lastLinePrinted := -1
	permitDelim := args.before > 0 || args.after > 0

	for window := range windowCh {
		beforeInWindow := min(lineNumber, args.before)
		line := window[beforeInWindow]
		match, err := matches(line, args)
		if err == nil && match {
			// output entire window
			trim := max(0, lastLinePrinted-lineNumber+beforeInWindow)
			needDelim := lineNumber-lastLinePrinted+1 >= args.before+args.after
			needDelim = needDelim && lastLinePrinted != -1
			err = outputWindow(w, window[trim:], needDelim && permitDelim)
			lastLinePrinted = lineNumber + 1 + args.after
		}

		if err != nil {
			return err
		}
		lineNumber++
	}

	return nil
}

func windowedReader(ctx context.Context, windowCh chan<- []string, r io.Reader, before, after int) {
	defer close(windowCh)
	maxSize := before + 1 + after
	var window []string
	sc := bufio.NewScanner(r)

	// scan after
	for i := 0; i < after && sc.Scan(); i++ {
		window = pushLine(window, sc.Text(), maxSize)
	}

	for sc.Scan() {
		if sc.Err() != nil || ctx.Err() != nil {
			return
		}
		window = pushLine(window, sc.Text(), maxSize)
		windowCh <- window
	}

	// handle last windows
	for before+1 < len(window) {
		window = popLine(window)
		windowCh <- window
	}
}

func outputWindow(w io.Writer, window []string, needDelim bool) error {
	if needDelim {
		_, err := w.Write([]byte("--\n"))
		if err != nil {
			return err
		}
	}
	for _, line := range window {
		if _, err := w.Write([]byte(line + "\n")); err != nil {
			return err
		}
	}
	return nil
}

func matches(line string, a args) (res bool, err error) {
	pattern := a.pattern

	if a.ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if a.fixed {
		res = strings.Contains(line, pattern)
	} else {
		res, err = regexp.MatchString(pattern, line)
	}

	if a.invert {
		res = !res
	}
	return res, err
}

func parseArgs(rawArgs []string) (args, error) {
	a := args{}

	var err error
	for err == nil && len(rawArgs) > 1 {
		first, rest := rawArgs[0], rawArgs[1:]
		switch first {
		case "-A":
			a.after, rest, err = parseNumericArg(rest)
		case "-B":
			a.before, rest, err = parseNumericArg(rest)
		case "-C":
			var ctx int
			ctx, rest, err = parseNumericArg(rest)
			a.before = ctx
			a.after = ctx
		case "-c":
			a.count, rest, err = parseNumericArg(rest)
		case "-i":
			a.ignoreCase = true
		case "-v":
			a.invert = true
		case "-F":
			a.fixed = true
		case "-n":
			a.lineNum = true
		default:
			err = errUnknownFlag
		}
		rawArgs = rest
	}

	if len(rawArgs) == 0 {
		err = errPatternExpected
	}

	if err != nil {
		return args{}, err
	}

	a.pattern = rawArgs[0]
	return a, nil
}

func parseNumericArg(args []string) (int, []string, error) {
	if len(args) == 0 {
		return 0, nil, errValueExpected
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, nil, err
	}
	return n, args[1:], nil
}

func pushLine(buffer []string, line string, maxSize int) []string {
	newBuffer := make([]string, len(buffer), maxSize)
	if len(buffer) < maxSize {
		copy(newBuffer, buffer)
		newBuffer = append(newBuffer, line)
	} else if len(buffer) == maxSize {
		copy(newBuffer[:maxSize-1], buffer[1:])
		newBuffer[maxSize-1] = line
	}
	return newBuffer
}

func popLine(buffer []string) []string {
	return buffer[1:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
