package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSortLines_NoArgs(t *testing.T) {
	commonLines := []string{
		"ghi 678",
		"21 a b",
		"def 45  Dec",
		"ghi 678",
		"abc  123   Jan",
		"3 cd",
	}

	cases := []struct {
		lines []string
		want  []string
	}{
		{
			lines: []string{
				"",
				"",
			},
			want: []string{
				"",
				"",
			},
		},
		{
			lines: commonLines,
			want: []string{
				"21 a b",
				"3 cd",
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
				"ghi 678",
			},
		},
	}

	for _, c := range cases {
		got, err := sortLines(args{}, c.lines)

		assert.NilError(t, err)
		assert.DeepEqual(t, got, c.want)
	}
}

func TestSortLines_Fields(t *testing.T) {
	commonLines := []string{
		"ghi 678",
		"21 a b",
		"def 45  Dec",
		"ghi 678",
		"abc  123   Jan",
		"3 cd",
	}

	cases := []struct {
		fields []int
		lines  []string
		want   []string
	}{
		{
			fields: nil,
			lines:  commonLines,
			want: []string{
				"21 a b",
				"3 cd",
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
				"ghi 678",
			},
		},
		{
			fields: []int{1},
			lines:  commonLines,
			want: []string{
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
				"ghi 678",
				"21 a b",
				"3 cd",
			},
		},
		{
			fields: []int{2, 1},
			lines:  commonLines,
			want: []string{
				"ghi 678",
				"ghi 678",
				"3 cd",
				"def 45  Dec",
				"abc  123   Jan",
				"21 a b",
			},
		},
	}

	for i, c := range cases {
		t.Log(i)
		got, err := sortLines(args{fields: c.fields}, c.lines)

		assert.NilError(t, err)
		assert.DeepEqual(t, got, c.want)
	}
}

func TestSortLines_Reverse(t *testing.T) {
	commonLines := []string{
		"ghi 678",
		"21 a b",
		"def 45  Dec",
		"ghi 678",
		"abc  123   Jan",
		"3 cd",
	}

	cases := []struct {
		reverse bool
		lines   []string
		want    []string
	}{
		{
			reverse: false,
			lines:   commonLines,
			want: []string{
				"21 a b",
				"3 cd",
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
				"ghi 678",
			},
		},
		{
			reverse: true,
			lines:   commonLines,
			want: []string{
				"ghi 678",
				"ghi 678",
				"def 45  Dec",
				"abc  123   Jan",
				"3 cd",
				"21 a b",
			},
		},
	}

	for i, c := range cases {
		t.Log(i)
		got, err := sortLines(args{reverse: c.reverse}, c.lines)

		assert.NilError(t, err)
		assert.DeepEqual(t, got, c.want)
	}
}

func TestSortLines_Unique(t *testing.T) {
	commonLines := []string{
		"ghi 678",
		"21 a b",
		"def 45  Dec",
		"ghi 678",
		"abc  123   Jan",
		"3 cd",
	}

	cases := []struct {
		unique bool
		lines  []string
		want   []string
	}{
		{
			unique: false,
			lines:  commonLines,
			want: []string{
				"21 a b",
				"3 cd",
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
				"ghi 678",
			},
		},
		{
			unique: true,
			lines:  commonLines,
			want: []string{
				"21 a b",
				"3 cd",
				"abc  123   Jan",
				"def 45  Dec",
				"ghi 678",
			},
		},
	}

	for i, c := range cases {
		t.Log(i)
		got, err := sortLines(args{unique: c.unique}, c.lines)

		assert.NilError(t, err)
		assert.DeepEqual(t, got, c.want)
	}
}
