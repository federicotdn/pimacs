package core

import (
	"testing"
)

func TestCharTableSet(t *testing.T) {
	t.Parallel()
	ec := newExecContext()

	type insertion struct {
		from, to lispInt
		text     string
	}

	type testCase struct {
		insertions []insertion
		expected   []lispCharTableEntry
	}

	cases := []testCase{
		{
			insertions: []insertion{},
			expected:   []lispCharTableEntry{},
		},
		{
			insertions: []insertion{
				{0, 1, "foo"},
			},
			expected: []lispCharTableEntry{
				{0, 1, ec.makeString("foo")},
			},
		},
		{
			insertions: []insertion{
				{0, 0, "foo"},
				{1, 1, "x"},
				{4, 4, "y"},
			},
			expected: []lispCharTableEntry{
				{0, 0, ec.makeString("foo")},
				{1, 1, ec.makeString("x")},
				{4, 4, ec.makeString("y")},
			},
		},
		{
			insertions: []insertion{
				{0, 0, "foo"},
				{1, 1, "x"},
				{4, 4, "y"},
				{0, 10, "--"},
			},
			expected: []lispCharTableEntry{
				{0, 10, ec.makeString("--")},
			},
		},
		{
			insertions: []insertion{
				{10, 20, "foo"},
				{15, 25, "foo"},
			},
			expected: []lispCharTableEntry{
				{10, 14, ec.makeString("foo")},
				{15, 25, ec.makeString("foo")},
			},
		},
		{
			insertions: []insertion{
				{0, 1, "foo"},
				{2, 10, "bar"},
			},
			expected: []lispCharTableEntry{
				{0, 1, ec.makeString("foo")},
				{2, 10, ec.makeString("bar")},
			},
		},
		{
			insertions: []insertion{
				{0, 1, "foo"},
				{1, 10, "bar"},
			},
			expected: []lispCharTableEntry{
				{0, 0, ec.makeString("foo")},
				{1, 10, ec.makeString("bar")},
			},
		},
		{
			insertions: []insertion{
				{0, 10, "foo"},
				{5, 15, "bar"},
			},
			expected: []lispCharTableEntry{
				{0, 4, ec.makeString("foo")},
				{5, 15, ec.makeString("bar")},
			},
		},
		{
			insertions: []insertion{
				{0, 10, "foo"},
				{5, 15, "bar"},
				{0, 15, "x"},
			},
			expected: []lispCharTableEntry{
				{0, 15, ec.makeString("x")},
			},
		},
		{
			insertions: []insertion{
				{0, 20, "foo"},
				{5, 15, "bar"},
			},
			expected: []lispCharTableEntry{
				{0, 4, ec.makeString("foo")},
				{5, 15, ec.makeString("bar")},
				{16, 20, ec.makeString("foo")},
			},
		},
		{
			insertions: []insertion{
				{0, maxChar, "foo"},
			},
			expected: []lispCharTableEntry{
				{0, maxChar, ec.makeString("foo")},
			},
		},
		{
			insertions: []insertion{
				{0, maxChar, "foo"},
				{10, 10, "bar"},
			},
			expected: []lispCharTableEntry{
				{0, 9, ec.makeString("foo")},
				{10, 10, ec.makeString("bar")},
				{11, maxChar, ec.makeString("foo")},
			},
		},
	}

	for i, case_ := range cases {
		ct := xCharTable(xEnsure(ec.makeCharTable(ec.nil_, ec.nil_)))
		for _, ins := range case_.insertions {
			ec.charTableSet(ct, ins.from, ins.to, ec.makeString(ins.text))
		}

		if len(ct.val) != len(case_.expected) {
			t.Logf("entries length mismatch: case index %v", i)
			t.Fail()
		}

		for j := 0; j < len(ct.val); j++ {
			have := ct.val[j]
			want := case_.expected[j]

			if have.end != want.end ||
				have.start != want.start ||
				xStringValue(have.val) != xStringValue(want.val) {

				t.Logf("entry mismatch: have '%+v', want '%+v'", have, want)
				t.Logf("entry mismatch: case index %v, entry %v", i, j)
				t.Fail()
			}
		}
	}
}

func TestCharTableLookup(t *testing.T) {
	t.Parallel()
	ec := newExecContext()
	ct := xCharTable(xEnsure(ec.makeCharTable(ec.nil_, ec.nil_)))

	ct.val = []lispCharTableEntry{
		{start: 2, end: 10},
		{start: 11, end: 20},
		{start: 30, end: 31},
	}

	type testCase struct {
		c             lispInt
		expectedIdx   int
		expectedFound bool
	}

	cases := []testCase{
		{c: 0, expectedIdx: 0, expectedFound: false},
		{c: 2, expectedIdx: 0, expectedFound: true},
		{c: 5, expectedIdx: 0, expectedFound: true},
		{c: 10, expectedIdx: 0, expectedFound: true},
		{c: 11, expectedIdx: 1, expectedFound: true},
		{c: 20, expectedIdx: 1, expectedFound: true},
		{c: 21, expectedIdx: 2, expectedFound: false},
		{c: 25, expectedIdx: 2, expectedFound: false},
		{c: 30, expectedIdx: 2, expectedFound: true},
		{c: 31, expectedIdx: 2, expectedFound: true},
		{c: 35, expectedIdx: 3, expectedFound: false},
		{c: 100, expectedIdx: 3, expectedFound: false},
	}

	for i, case_ := range cases {
		idx, found := ec.charTableLookup(ct, case_.c)
		if idx != case_.expectedIdx || found != case_.expectedFound {
			t.Logf("mismatched character table lookup: case index %v", i)
			t.Fail()
		}
	}
}
