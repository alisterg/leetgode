package easy

import (
	"leetgode/test"
	"reflect"
	"testing"
)

func TestArrayDiffExampleOne(t *testing.T) {
	input := []int{1, 2, 3}
	input2 := []int{2, 4, 6}
	expected := [][]int{{1, 3}, {4, 6}}
	actual := arrayDiff(input, input2)

	if !reflect.DeepEqual(actual, expected) {
		test.Fail(t, expected, actual)
	}
}

func TestArrayDiffExampleTwo(t *testing.T) {
	input := []int{1, 2, 3, 3}
	input2 := []int{1, 1, 2, 2}
	expected := [][]int{{3}, {}}
	actual := arrayDiff(input, input2)

	if !reflect.DeepEqual(actual, expected) {
		test.Fail(t, expected, actual)
	}
}
