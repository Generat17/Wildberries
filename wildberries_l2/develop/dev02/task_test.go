package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	cases := []struct {
		input, expected string
		err             error
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			err:      nil,
		},
		{
			input:    "abcd",
			expected: "abcd",
			err:      nil,
		},
		{
			input:    "45",
			expected: "",
			err:      errInvalidFormat,
		},
		{
			input:    "",
			expected: "",
			err:      nil,
		},
	}

	for _, c := range cases {
		actual, err := UnpackString(c.input)

		if c.err == nil {
			assert.NilError(t, err)
		} else {
			assert.ErrorIs(t, err, c.err)
		}
		assert.Equal(t, c.expected, actual)
	}
}

func TestUnpackEscapeString(t *testing.T) {
	cases := []struct {
		input, expected string
		err             error
	}{
		{
			input:    `qwe\4\5`,
			expected: "qwe45",
			err:      nil,
		},
		{
			input:    `qwe\45`,
			expected: "qwe44444",
			err:      nil,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
			err:      nil,
		},
		{
			input:    `qwe\\5\`,
			expected: "",
			err:      errInvalidEscapeSequence,
		},
	}

	for _, c := range cases {
		actual, err := UnpackString(c.input)

		if c.err == nil {
			assert.NilError(t, err)
		} else {
			assert.ErrorIs(t, err, c.err)
		}
		assert.Equal(t, c.expected, actual)
	}
}
