package unoptimized_test

import (
	"strconv"
	"testing"

	"github.com/jboursiquot/go-for-experienced-programmers/04-testing/benchmarking/wordlens"
	"github.com/jboursiquot/go-for-experienced-programmers/04-testing/benchmarking/wordlens/unoptimized"
)

func TestFindPalindromes(t *testing.T) {
	tests := map[string]struct {
		words    []string
		expected int
	}{
		"empty": {
			words:    []string{},
			expected: 0,
		},
		"single": {
			words:    []string{"a"},
			expected: 1,
		},
		"multiple": {
			words:    []string{"racecar", "hello", "madam", "level", "deified", "rotator", "step on no pets", "not a palindrome"},
			expected: 6,
		},
		"lots": {
			words:    wordlens.TestWords(),
			expected: 27,
		},
	}

	wl := unoptimized.NewWordLens()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res := wl.FindPalindromes(tc.words)
			if len(res) != tc.expected {
				t.Errorf("Expected %d palindromes, got %d", tc.expected, len(res))
			}
		})
	}
}

func BenchmarkFindPalindromes(b *testing.B) {
	b.StopTimer() // exclude preparations from the benchmark
	wl := unoptimized.NewWordLens()
	allWords := wordlens.TestWords()

	b.StartTimer() // run the benchmark
	for n := 25; n <= len(allWords); n = n + 25 {
		words := allWords[:n]
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wl.FindPalindromes(words)
			}
		})
	}
}
