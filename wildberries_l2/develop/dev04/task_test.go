package main

import (
	"reflect"
	"testing"
)

func TestNoAnagrams(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s:    nil,
			want: map[string][]string{},
		},
		{
			s:    []string{"первый", "второй"},
			want: map[string][]string{},
		},
	}

	for _, c := range cases {
		if got := extractAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("extractAnagrams(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}

func TestAnagramsLowerCase(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s: []string{"строка", "пятак", "пятка", "тяпка", "листок", "сорока", "слиток", "столик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, c := range cases {
		if got := extractAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("extractAnagrams(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}

func TestAnagramsBothCase(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s: []string{"строка", "ПяТак", "пЯткА", "ТЯПка", "лИСток", "сОвОка", "слИТОк", "сТОлик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, c := range cases {
		if got := extractAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("extractAnagrams(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}
