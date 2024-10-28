package easy

import (
	"unicode"
)

// Given a string s, reverse only all the vowels in the string and return the
// result.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower
// and upper cases, more than once.
//
// Example: given "IceCreAm", produces "AceCreIm"
// Example: given "leetcode", produces "leotcede"
func reverse(s string) string {
	foundVowels := ""
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			foundVowels += string(s[i])
		}
	}

	result := ""
	for i := 0; i < len(s); i++ {
		if !isVowel(s[i]) {
			result += string(s[i])
			continue
		}

		lastVowel := foundVowels[len(foundVowels)-1]
		foundVowels = foundVowels[:len(foundVowels)-1]
		result += string(lastVowel)
	}

	return result
}

func isVowel(b byte) bool {
	c := unicode.ToLower(rune(b))
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
