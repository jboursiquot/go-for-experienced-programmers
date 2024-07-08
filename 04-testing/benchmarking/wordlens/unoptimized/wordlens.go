package unoptimized

type WordLens struct{}

func NewWordLens() WordLens {
	return WordLens{}
}

func (wl WordLens) FindPalindromes(words []string) map[string]int {
	res := make(map[string]int)
	for _, word := range words {
		if wl.isPalindrome(word) {
			res[word]++
		}
	}

	return res
}

func (wl WordLens) isPalindrome(word string) bool {
	runes := []rune(word)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}
