package main

import "testing"

func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 1},
		"two":   {input: "812 %6ab//", want: 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "abc 1,?/", want: 1},
		"two":   {input: "abc 12&8 ^", want: 3},
	}

	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := nc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestSymbolCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "abc 1,?/", want: 4},
		"two":   {input: "abc 12&8 ^_", want: 5},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := sc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
