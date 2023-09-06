package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

# Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	errUnknownArgument   = errors.New("unknown argument")
	errFieldsNotProvided = errors.New("fields not provided")
)

type args struct {
	Fields    []int
	Delim     string
	Separated bool
}

func main() {
	args, err := parseArgs(os.Args[1:])
	if err == nil {
		err = applyCut(os.Stdout, os.Stdin, args)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func applyCut(out io.Writer, in io.Reader, cmdArgs args) error {
	var err error

	for sc := bufio.NewScanner(in); sc.Scan() && err == nil; {
		line := sc.Text()
		if strings.Contains(line, cmdArgs.Delim) {
			cells := strings.Split(line, cmdArgs.Delim)
			err = printOnlyFields(out, cells, cmdArgs.Fields, cmdArgs.Delim)
		} else if !cmdArgs.Separated {
			_, err = out.Write([]byte(line + "\n"))
		}
	}

	return err
}

func printOnlyFields(out io.Writer, strs []string, fields []int, delim string) error {
	sb := strings.Builder{}

	notFirst := false
	for _, i := range fields {
		if notFirst {
			sb.WriteString(delim)
		}
		notFirst = true
		if i < len(strs) {
			sb.WriteString(strs[i])
		}
	}
	sb.WriteByte('\n')

	_, err := out.Write([]byte(sb.String()))
	return err
}

func parseArgs(rawArgs []string) (args, error) {
	Fields := []int{}
	Delim := "\t"
	Separated := false

	for i := 0; i < len(rawArgs); i++ {
		var err error

		switch rawArgs[i] {
		case "-f":
			if i++; i >= len(rawArgs) {
				err = errFieldsNotProvided
			} else {
				Fields, err = parseFields(rawArgs[i])
			}
		case "-d":
			if i++; i >= len(rawArgs) {
				err = errFieldsNotProvided
			} else {
				Delim = rawArgs[i]
			}
		case "-s":
			Separated = true
		default:
			err = errUnknownArgument
		}

		if err != nil {
			return args{}, err
		}
	}

	if len(Fields) == 0 {
		return args{}, errFieldsNotProvided
	}

	return args{
		Fields:    Fields,
		Delim:     Delim,
		Separated: Separated,
	}, nil
}

func parseFields(args string) ([]int, error) {
	fieldsStr := strings.Split(args, ",")
	fields := make([]int, len(fieldsStr))
	for i := 0; i < len(fieldsStr); i++ {
		var err error
		fields[i], err = strconv.Atoi(fieldsStr[i])
		if err != nil {
			return nil, err
		}
	}
	return fields, nil
}
