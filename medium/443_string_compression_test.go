package medium

import (
	"leetgode/test"
	"reflect"
	"testing"
)

func TestCompressExampleOne(t *testing.T) {
	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	expected := []byte{'a', '2', 'b', '2', 'c', '3', 'c'}

	compress(input)

	if !reflect.DeepEqual(input, expected) {
		test.Fail(t, expected, input)
	}
}

func TestCompressExampleTwo(t *testing.T) {
	input := []byte{'a'}
	expected := []byte{'a'}

	compress(input)

	if !reflect.DeepEqual(input, expected) {
		test.Fail(t, expected, input)
	}
}

func TestCompressExampleThree(t *testing.T) {
	input := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	expected := []byte{'a', 'b', '1', '2', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}

	compress(input)

	if !reflect.DeepEqual(input, expected) {
		test.Fail(t, expected, input)
	}
}
