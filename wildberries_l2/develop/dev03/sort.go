package main

import (
	"sort"
	"strconv"
	"strings"
)

type cmp int

const (
	less cmp = iota - 1
	eq
	greater
)

func sortLines(args args, lines []string) ([]string, error) {
	sorterFunc := buildSorter(args.fields, args.numeric)

	sort.SliceStable(lines, func(i, j int) bool {
		return sorterFunc(lines[i], lines[j]) != greater
	})

	if args.reverse {
		reverse(lines)
	}

	if args.unique {
		filterUnique(&lines)
	}

	return lines, nil
}

func buildSorter(fields []int, numeric bool) func(string, string) cmp {
	sorter := strCmp
	if numeric {
		sorter = numCmp
	}

	if len(fields) == 0 {
		return sorter
	}

	return func(s1, s2 string) cmp {
		s1f := extractFields(s1, fields)
		s2f := extractFields(s2, fields)

		for i := 0; i < len(fields); i++ {
			cmp := sorter(s1f[i], s2f[i])
			if cmp != eq {
				return cmp
			}
		}

		return eq
	}
}

func strCmp(s1, s2 string) cmp {
	if s1 < s2 {
		return less
	} else if s1 > s2 {
		return greater
	}
	return eq
}

func numCmp(s1, s2 string) cmp {
	n1, err1 := strconv.Atoi(s1)
	n2, err2 := strconv.Atoi(s2)
	if err1 != nil || err2 != nil {
		return strCmp(s1, s2)
	}
	if n1 < n2 {
		return less
	} else if n1 == n2 {
		return eq
	} else {
		return greater
	}
}

func extractFields(s string, fields []int) []string {
	sf := strings.Fields(s)
	res := make([]string, len(fields))

	for i, field := range fields {
		if field < len(sf) {
			res[i] = sf[field]
		}
	}

	return res
}

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func filterUnique(lines *[]string) {
	var i int
	for j := 1; j < len(*lines); j++ {
		if (*lines)[i] != (*lines)[j] {
			i++
			(*lines)[i] = (*lines)[j]
		}
	}
	*lines = (*lines)[:i+1]
}
