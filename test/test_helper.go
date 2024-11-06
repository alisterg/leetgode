package test

import (
	"testing"
)

func Fail(t *testing.T, expected interface{}, actual interface{}) {
	t.Errorf("Expected `%v`, got `%v`", expected, actual)
}
