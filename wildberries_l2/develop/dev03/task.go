package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// read cmdline args
	args, err := parseArgs(normalizeArgs(os.Args[1:]))
	if err != nil {
		log.Fatal(err)
	}

	lines, err := readLines(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines, err = sortLines(args, lines)
	if err != nil {
		log.Fatal(err)
	}

	err = writeLines(os.Stdout, lines)
	if err != nil {
		log.Fatal(err)
	}
}

func readLines(r io.Reader) ([]string, error) {
	scan := bufio.NewScanner(r)
	lines := make([]string, 0)
	for scan.Scan(); scan.Err() == nil; scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, scan.Err()
}

func writeLines(w io.Writer, lines []string) error {
	sb := strings.Builder{}
	for _, line := range lines {
		sb.WriteString(line)
		sb.WriteByte('\n')
	}
	sb.WriteByte('\n')
	_, err := io.WriteString(w, sb.String())
	return err
}
