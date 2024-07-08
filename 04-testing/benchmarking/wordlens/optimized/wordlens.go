package optimized

type WordLens struct{}

func NewWordLens() WordLens {
	return WordLens{}
}

func (wl WordLens) FindPalindromes(words []string) map[string]int {
	palindromes := make(map[string]int)
	for _, word := range words {
		if wl.isPalindrome(word) {
			palindromes[word]++
		}
	}
	return palindromes
}

func (wl WordLens) isPalindrome(word string) bool {
	for i := range word {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}
