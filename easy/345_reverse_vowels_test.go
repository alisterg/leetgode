package easy

import (
	"leetgode/test"
	"testing"
)

func TestReverseExampleOne(t *testing.T) {
	input := "IceCreAm"
	expected := "AceCreIm"

	if actual := reverseVowels(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestReverseExampleTwo(t *testing.T) {
	input := "leetcode"
	expected := "leotcede"

	if actual := reverseVowels(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}
