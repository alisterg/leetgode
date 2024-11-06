package easy

import (
	"leetgode/test"
	"reflect"
	"testing"
)

func TestMoveZeroesExampleOne(t *testing.T) {
	input := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(input)

	if !reflect.DeepEqual(input, expected) {
		test.Fail(t, expected, input)
	}
}

func TestMoveZeroesExampleTwo(t *testing.T) {
	input := []int{0}
	expected := []int{0}
	moveZeroes(input)

	if !reflect.DeepEqual(input, expected) {
		test.Fail(t, expected, input)
	}
}
