package core

import (
	"encoding/json"
	"os"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Parallel()
	ec := newTestingInterpreter().ec

	cases := [][]lispObject{
		{},
		{newInteger(1)},
		{newInteger(123), newInteger(2), newInteger(10)},
	}

	for _, c := range cases {
		head := ec.makeList(c...)
		produced := []lispObject{}

		iter := ec.iterate(head)
		for ; iter.hasNext(); head = iter.nextCons() {
			produced = append(produced, xCar(head))
		}

		if iter.hasError() {
			t.Fail()
		}

		if len(c) != len(produced) {
			t.Fail()
		}

		for i := 0; i < len(c); i++ {
			if xIntegerValue(c[i]) != xIntegerValue(produced[i]) {
				t.Fail()
			}
		}
	}
}

func TestIterationFail(t *testing.T) {
	t.Parallel()
	ec := newTestingInterpreter().ec

	heads := []lispObject{
		newCons(newInteger(1), newInteger(10)),
		newCons(newInteger(1), newCons(newInteger(1), newInteger(10))),
		newInteger(100),
	}

	for _, head := range heads {
		iter := ec.iterate(head)
		for ; iter.hasNext(); iter.nextCons() {
		}

		if !iter.hasError() {
			t.Fail()
		}
	}
}

type subroutine struct {
	Lname   string `json:"lname"`
	MinArgs int    `json:"minargs"`
	MaxArgs int    `json:"maxargs"`
}

type subroutineData struct {
	Subroutines []subroutine `json:"subroutines"`
}

func TestSubroutineSignatures(t *testing.T) {
	t.Parallel()
	ec := newTestingInterpreter().ec

	data, err := os.ReadFile("../test/data/emacs_subroutines.json")
	if err != nil {
		t.Fatal(err)
	}

	var sd subroutineData
	err = json.Unmarshal(data, &sd)
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range sd.Subroutines {
		sym, ok := ec.obarray.val[s.Lname]
		if !ok || !subroutinep(sym.function) {
			continue
		}

		subr := xSubroutine(sym.function)

		if subr.minArgs != s.MinArgs {
			t.Errorf("minargs mismatch for: '%v'", s.Lname)
			continue
		}

		if subr.maxArgs != s.MaxArgs {
			t.Errorf("maxargs mismatch for: '%v'", s.Lname)
		}
	}
}
