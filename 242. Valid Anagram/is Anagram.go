package _42__Valid_Anagram

import "strings"

func isAnagram(s string, t string) bool {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")
	t = strings.ReplaceAll(strings.ToLower(t), " ", "")

	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for _, char := range t {
		if count, found := charCount[char]; !found || count == 0 {
			return false
		}
		charCount[char]--
	}

	return true
}
