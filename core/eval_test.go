package core

import (
	"testing"
)

func TestEvalLisp(t *testing.T) {
	t.Parallel()
	testLisp(t, "eval_test.el")
}
