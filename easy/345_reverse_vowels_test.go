package easy

import (
	"leetgode/test"
	"testing"
)

func TestReverseExampleOne(t *testing.T) {
	input := "IceCreAm"
	expected := "AceCreIm"

	if actual := reverse(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestReverseExampleTwo(t *testing.T) {
	input := "leetcode"
	expected := "leotcede"

	if actual := reverse(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}
