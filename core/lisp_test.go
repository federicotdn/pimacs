package core

import (
	"strings"
	"testing"
)

var filenames = []string{
	"character-table-tests.el",
	"eval-tests.el",
	"functions-tests.el",
	"goroutine-tests.el",
	"lisp-tests.el",
	"read-tests.el",
}

func TestLisp(t *testing.T) {
	t.Parallel()
	inp := newTestingInterpreter()
	err := inp.LoadFile("lt.el")
	if err != nil {
		t.Logf("failed to load 'lt.el': %v", err)
		t.FailNow()
	}

	for _, filename := range filenames {
		err = inp.LoadFile(filename)
		if err != nil {
			t.Logf("failed to load '%v': %+v", filename, err)
			t.FailNow()
		}
		_, err = inp.ReadEvalPrin1("(lt--run-all-tests)")
		if err != nil {
			t.Logf("test(s) failure in '%v': %+v", filename, strings.Split(err.Error(), "\n")[0])
			info, err2 := inp.ReadEvalPrin1("lt--failure-info")
			if err2 != nil {
				t.Logf("failed to retrieve failure info: %+v", err2)
			} else {
				t.Logf("test(s) failure info: %+v", info)
			}
			t.Fail()
		}
	}
}
