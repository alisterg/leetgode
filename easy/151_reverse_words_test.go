package easy

import (
	"leetgode/test"
	"testing"
)

func TestReverseWordsExampleOne(t *testing.T) {
	input := "the sky is blue"
	expected := "blue is sky the"

	if actual := reverseWords(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestReverseWordsExampleTwo(t *testing.T) {
	input := "  hello world  "
	expected := "world hello"

	if actual := reverseWords(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestReverseWordsExampleThree(t *testing.T) {
	input := "a good   example"
	expected := "example good a"

	if actual := reverseWords(input); actual != expected {
		test.Fail(t, expected, actual)
	}
}
