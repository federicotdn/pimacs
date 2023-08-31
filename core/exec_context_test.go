package core

import (
	"testing"
)

func TestIteration(t *testing.T) {
	t.Parallel()
	ec := newExecContext()

	cases := [][]lispObject{
		{},
		{ec.makeInteger(1)},
		{ec.makeInteger(123), ec.makeInteger(2), ec.makeInteger(10)},
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
	ec := newExecContext()

	heads := []lispObject{
		ec.makeCons(ec.makeInteger(1), ec.makeInteger(10)),
		ec.makeCons(ec.makeInteger(1), ec.makeCons(ec.makeInteger(1), ec.makeInteger(10))),
		ec.makeInteger(100),
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
