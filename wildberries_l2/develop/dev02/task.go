package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	errInvalidFormat         = errors.New("invalid packed format")
	errInvalidEscapeSequence = errors.New("invalid escape sequence")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := scanner.Text()
	s, err := UnpackString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}

func UnpackString(s string) (string, error) {
	runes := []rune(s)
	sb := strings.Builder{}

	for pos := 0; pos < len(runes); pos++ {
		if unicode.IsDigit(runes[pos]) {
			return "", errInvalidFormat
		}

		// обработка escape-последовательности
		if runes[pos] == '\\' {
			pos++
			if pos == len(runes) {
				return "", errInvalidEscapeSequence
			}
		}

		// если последняя руна или следующая не цифра
		if pos+1 == len(runes) || !unicode.IsDigit(runes[pos+1]) {
			sb.WriteRune(runes[pos])
			continue
		}

		// дубликат руны на позиции
		multiplier := int(runes[pos+1] - '0')
		nextPos := pos + 2
		for nextPos < len(runes) && unicode.IsDigit(runes[nextPos]) {
			multiplier = 10*multiplier + int(runes[nextPos]-'0')
			nextPos++
		}
		sb.Grow(multiplier * utf8.RuneLen(runes[pos]))
		for i := 0; i < multiplier; i++ {
			sb.WriteRune(runes[pos])
		}
		pos = nextPos - 1
	}

	return sb.String(), nil
}
