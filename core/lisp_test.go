package core

import (
	"strings"
	"testing"
)

func testLisp(t *testing.T, filename string) {
	inp := newTestingInterpreter()
	err := inp.LoadFile("lisp_test.el")
	if err != nil {
		t.Logf("failed to load 'lisp_test.el': %v", err)
		t.Fail()
	}
	err = inp.LoadFile(filename)
	if err != nil {
		t.Logf("failed to load '%v': %+v", filename, err)
		t.Fail()
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
