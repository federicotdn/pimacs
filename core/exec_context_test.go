package core

import (
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
		for ; iter.hasNext(); head = iter.nextCons() {
		}

		if !iter.hasError() {
			t.Fail()
		}
	}
}
