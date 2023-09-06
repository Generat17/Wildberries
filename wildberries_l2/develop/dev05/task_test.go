package main

import (
	"log"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGrep_NoContext(t *testing.T) {
	cases := []struct {
		pattern string
		lines   []string
		want    []string
	}{
		{
			pattern: "a",
			lines: []string{
				"line 1 a",
				"line 2",
				"line 3 a",
				"line 4 a",
			},
			want: []string{
				"line 1 a",
				"line 3 a",
				"line 4 a",
			},
		},
	}

	for _, c := range cases {
		r := strings.NewReader(strings.Join(c.lines, "\n") + "\n")
		w := strings.Builder{}

		err := runGrep(&w, r, args{pattern: c.pattern})
		assert.NilError(t, err)

		got := w.String()
		want := strings.Join(c.want, "\n") + "\n"
		if got != want {
			t.Errorf("grep(%v) = %v, want %v", c.lines, got, c.want)
		}
	}
}

func TestGrep_WithContext(t *testing.T) {
	lines := []string{
		"line 1 a",
		"line 2",
		"line 3",
		"line 4 a",
		"line 5",
		"line 6",
		"line 7",
		"line 8",
		"line 9 a",
		"line 10",
	}
	cases := []struct {
		pattern string
		before  int
		after   int
		lines   []string
		want    []string
	}{
		{
			pattern: "a",
			before:  1,
			after:   1,
			lines:   lines,
			want: []string{
				"line 1 a",
				"line 2",
				"--",
				"line 3",
				"line 4 a",
				"line 5",
				"--",
				"line 8",
				"line 9 a",
				"line 10",
			},
		},
		{
			pattern: "a",
			before:  1,
			after:   2,
			lines:   lines,
			want: []string{
				"line 1 a",
				"line 2",
				"line 3",
				"line 4 a",
				"line 5",
				"line 6",
				"--",
				"line 8",
				"line 9 a",
				"line 10",
			},
		},
		{
			pattern: "a",
			before:  2,
			after:   2,
			lines:   lines,
			want: []string{
				"line 1 a",
				"line 2",
				"line 3",
				"line 4 a",
				"line 5",
				"line 6",
				"line 7",
				"line 8",
				"line 9 a",
				"line 10",
			},
		},
		{
			pattern: "a",
			before:  1,
			after:   0,
			lines:   lines,
			want: []string{
				"line 1 a",
				"--",
				"line 3",
				"line 4 a",
				"--",
				"line 8",
				"line 9 a",
			},
		},
		{
			pattern: "a",
			before:  0,
			after:   1,
			lines:   lines,
			want: []string{
				"line 1 a",
				"line 2",
				"--",
				"line 4 a",
				"line 5",
				"--",
				"line 9 a",
				"line 10",
			},
		},
		{
			pattern: "a",
			before:  0,
			after:   2,
			lines:   lines,
			want: []string{
				"line 1 a",
				"line 2",
				"line 3",
				"line 4 a",
				"line 5",
				"line 6",
				"--",
				"line 9 a",
				"line 10",
			},
		},
	}

	for i, c := range cases {
		log.Println(i)
		r := strings.NewReader(strings.Join(c.lines, "\n") + "\n")
		w := strings.Builder{}

		err := runGrep(&w, r, args{pattern: c.pattern, before: c.before, after: c.after})
		assert.NilError(t, err)

		got := w.String()
		want := strings.Join(c.want, "\n") + "\n"
		if got != want {
			t.Errorf("grep(%v) = %v, want %v", c.lines, got, c.want)
		}
	}
}
