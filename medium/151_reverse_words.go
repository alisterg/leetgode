package medium

import (
	"strings"
)

// Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will
// be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
// Note that s may contain leading or trailing spaces or multiple spaces between
// two words. The returned string should only have a single space separating the
// words. Do not include any extra spaces.
//
// `s` contains alphanumeric characters and spaces only.
func reverseWords(s string) string {
	split := strings.Split(s, " ")
	var filtered []string

	for i := range split {
		if strings.TrimSpace(split[i]) != "" {
			filtered = append(filtered, split[i])
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		filtered[i], filtered[j] = strings.TrimSpace(filtered[j]), strings.TrimSpace(filtered[i])
	}

	return strings.Join(filtered, " ")
}

/*

const reverse = s => s
	.split(" ")
	.filter(el => el.trim() !== "")
	.reverse()
	.join(" ")

In JS...

*/
