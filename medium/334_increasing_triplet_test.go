package medium

import (
	"leetgode/test"
	"testing"
)

func TestIncreasingTripletExampleOne(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := true

	result := increasingTriplet(input)

	if result != expected {
		test.Fail(t, expected, result)
	}
}

func TestIncreasingTripletExampleTwo(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	expected := false

	result := increasingTriplet(input)

	if result != expected {
		test.Fail(t, expected, result)
	}
}

func TestIncreasingTripletExampleThree(t *testing.T) {
	input := []int{2, 1, 5, 0, 4, 6}
	expected := true

	result := increasingTriplet(input)

	if result != expected {
		test.Fail(t, expected, result)
	}
}

func TestIncreasingTripletExampleFour(t *testing.T) {
	input := []int{20, 100, 10, 12, 5, 13}
	expected := true

	result := increasingTriplet(input)

	if result != expected {
		test.Fail(t, expected, result)
	}
}
