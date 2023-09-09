package core

import (
	"testing"
)

func TestGoroutineLisp(t *testing.T) {
	t.Parallel()
	testLisp(t, "goroutine_test.el")
}
