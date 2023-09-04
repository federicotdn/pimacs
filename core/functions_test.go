package core

import (
	"testing"
)

func TestFunctionsLisp(t *testing.T) {
	t.Parallel()
	testLisp(t, "functions_test.el")
}
