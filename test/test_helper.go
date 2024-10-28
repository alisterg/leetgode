package test

import (
	"testing"
)

func Fail(t *testing.T, expected string, actual string) {
	t.Errorf("Expected `%v`, got `%v`", expected, actual)
}
