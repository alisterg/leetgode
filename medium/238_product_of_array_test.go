package medium

import (
	"leetgode/test"
	"reflect"
	"testing"
)

func TestProductOfArrayExampleOne(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	result := productExceptSelf(input)

	if !reflect.DeepEqual(result, expected) {
		test.Fail(t, expected, result)
	}
}

func TestProductOfArrayExampleTwo(t *testing.T) {
	input := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	result := productExceptSelf(input)

	if !reflect.DeepEqual(result, expected) {
		test.Fail(t, expected, result)
	}
}
