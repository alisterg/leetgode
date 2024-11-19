package easy

import (
	"leetgode/test"
	"testing"
)

func TestGcdOfStringsExampleOne(t *testing.T) {
	input := "ABCABC"
	input2 := "ABC"
	expected := "ABC"

	if actual := gcdOfStrings(input, input2); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestGcdOfStringsExampleTwo(t *testing.T) {
	input := "ABABAB"
	input2 := "ABAB"
	expected := "AB"

	if actual := gcdOfStrings(input, input2); actual != expected {
		test.Fail(t, expected, actual)
	}
}

func TestGcdOfStringsExampleFour(t *testing.T) {
	input := "LEET"
	input2 := "CODE"
	expected := ""

	if actual := gcdOfStrings(input, input2); actual != expected {
		test.Fail(t, expected, actual)
	}
}
