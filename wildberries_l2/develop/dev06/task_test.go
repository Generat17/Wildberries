package main

import (
	"strings"
	"testing"
)

func TestApplyCut(t *testing.T) {
	cases := []struct {
		args []string
		in   string
		want string
	}{
		{
			args: []string{"-f", "1,2,3"},
			in:   "",
			want: "",
		},
		{
			args: []string{"-f", "0", "-s"},
			in:   "abc\ndef\n",
			want: "",
		},
		{
			args: []string{"-f", "0", "-s"},
			in:   "a\tb\tc\nd e f\n",
			want: "a\n",
		},
		{
			args: []string{"-f", "2,0", "-s"},
			in:   "a\tb\tc\nd e f\n",
			want: "c\ta\n",
		},
		{
			args: []string{"-f", "1,1,1", "-s"},
			in:   "a\tb\tc\nd e f\n",
			want: "b\tb\tb\n",
		},
		{
			args: []string{"-f", "0,1,2", "-s", "-d", " "},
			in:   "a\tb\tc\nd e f\n",
			want: "d e f\n",
		},
	}

	for _, c := range cases {
		args, err := parseArgs(c.args)
		in := strings.NewReader(c.in)
		out := strings.Builder{}
		if err != nil {
			t.Errorf("parseArgs(%q) unexpected error: %v", c.args, err)
		}

		err = applyCut(&out, in, args)
		if err != nil {
			t.Errorf("applyCut(%q, %q) unexpected error: %v", c.args, c.in, err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("applyCut(%q) = %v, want %v", c.args, got, c.want)
		}
	}
}
