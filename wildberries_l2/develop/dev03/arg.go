package main

import (
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	errUnknownArgument   = errors.New("unknown argument")
	errNoColumnsProvided = errors.New("no columns provided")
)

type args struct {
	fields  []int
	numeric bool
	reverse bool
	unique  bool
}

func parseArgs(rawArgs []string) (args, error) {
	var fields []int
	var numeric, reverse, unique bool
	var err error

	for err == nil && len(rawArgs) > 0 {
		first, rest := rawArgs[0], rawArgs[1:]
		switch first {
		case "-k":
			fields, rest, err = parseFields(rest)
		case "-n":
			numeric = true
		case "-r":
			reverse = true
		case "-u":
			unique = true
		default:
			return args{}, errUnknownArgument
		}
		rawArgs = rest
	}

	if err != nil {
		return args{}, err
	}

	return args{fields, numeric, reverse, unique}, nil
}

func parseFields(rawArgs []string) ([]int, []string, error) {
	var fields []int

	for len(rawArgs) > 0 {
		field, err := strconv.Atoi(rawArgs[0])
		if err != nil {
			break
		}
		fields = append(fields, field)
		rawArgs = rawArgs[1:]
	}

	if len(fields) == 0 {
		return nil, nil, errNoColumnsProvided
	}

	return fields, rawArgs, nil
}

// splits args like '-ru' to '-r' and '-u'.
// Also splits keys from arguments like '-k2' to '-k' and '2'.
// Also filters empty strings in arguments.
func normalizeArgs(rawArgs []string) []string {
	result := make([]string, 0, len(rawArgs))

	for _, rawArg := range rawArgs {
		if len(rawArg) == 0 {
			continue // skip empty argument strings
		}

		if rawArg[0] != '-' { // found regular argument - leave it as it is
			result = append(result, rawArg)
			continue
		}

		// found key group
		if len(rawArgs) == 2 { // does not need to be splitted
			result = append(result, rawArg)
		} else {
			// split required
			result = append(result, splitArg(rawArg)...)
		}
	}

	return result
}

// splits args like '-ru' to '-r' and '-u'.
// Also splits keys from arguments like '-k2' to '-k' and '2'
func splitArg(argGroup string) []string {
	result := make([]string, 0, 1)

	sc := scanner.Scanner{}
	sc.Init(strings.NewReader(argGroup[1:]))
	sc.Mode = scanner.ScanChars | scanner.ScanInts

	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		if tok == scanner.Int {
			result = append(result, sc.TokenText())
		} else {
			result = append(result, "-"+sc.TokenText())
		}
	}

	return result
}
