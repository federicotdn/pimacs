package core

import (
	"testing"
)

func TestReadLisp(t *testing.T) {
	t.Parallel()
	testLisp(t, "read_test.el")
}
