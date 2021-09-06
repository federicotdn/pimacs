package elisp

import (
	"fmt"
	"testing"
)

type readStringTestCase struct {
	input    string
	expected string
	err      error
}

var anyError = fmt.Errorf("anyError")

func testReadStringHelper(t *testing.T, cases []readStringTestCase) {
	for _, tc := range cases {
		if tc.err != nil && tc.expected != "" {
			t.Fatalf("input: '%v' defines both expected value and error", tc.input)
		}

		interpreter := NewInterpreter()

		output, err := interpreter.ReadString(tc.input)
		if err != nil {
			if tc.err == nil {
				t.Logf("input: '%v' got error '%v'", tc.input, err)
				t.Fail()
			} else if tc.err != anyError && tc.err != err {
				t.Logf("input: '%v' got error '%v' instead of error '%v'", tc.input, err, tc.err)
				t.Fail()
			}
		} else {
			printed, err := interpreter.Print(output)
			if err != nil {
				t.Fatalf("input: '%v' unable to print output", tc.input)
			}

			if tc.err != nil {
				t.Logf("input: '%v' expected error '%v' instead of value '%v'", tc.input, tc.err, printed)
				t.Fail()
			} else if printed != tc.expected {
				t.Logf("input: '%v' expected '%v' but got '%v'", tc.input, tc.expected, printed)
				t.Fail()
			}
		}
	}
}

func TestReadStringInteger(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"1", "1", nil},
		{"-1", "-1", nil},
		{"0", "0", nil},
		{"-0", "0", nil},
	})
}

func TestReadStringFloat(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"1.2", "1.2", nil},
		{"0.111", "0.111", nil},
		{".11", "0.11", nil},
		{"123.123e-2", "1.23123", nil},
	})
}

func TestReadStringNil(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"nil", "nil", nil},
		{"()", "nil", nil},
	})
}

func TestReadStringCons(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"(1 . 1)", "(1 . 1)", nil},
		{"( . 2)", "2", nil},
		{"(. 2)", "2", nil},
		{"(. .)", ".", nil},
		{"(1 . (2 . 3))", "(1 2 . 3)", nil},
	})
}

func TestReadStringList(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"(1 1)", "(1 1)", nil},
		{"( 1 1 )", "(1 1)", nil},
		{"(.1 2)", "(0.1 2)", nil},
		{"(.)", "(.)", nil},
		{"((1 1))", "((1 1))", nil},
		{"((1 2 3) (4 5 6))", "((1 2 3) (4 5 6))", nil},
		{"(1 . (2 . (3 . nil)))", "(1 2 3)", nil},
		{"(", "", anyError},
		{"(123", "", anyError},
		{"(123 1 2 3", "", anyError},
		{"(123 1 2 3(", "", anyError},
		{"((())", "", anyError},
		{")", "", anyError},
		{"())", "nil", nil},
	})
}

func TestReadStringSymbol(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"foo", "foo", nil},
		{"foo\\ bar", "foo bar", nil},
		{"foo:bar", "foo:bar", nil},
		{"foo-bar", "foo-bar", nil},
		{"foo-bar?", "foo-bar?", nil},
		{"\\99", "99", nil},
		{"\\", "", anyError},
		{"foo\\", "", anyError},
	})
}

func TestReadStringQuote(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"'(1 1)", "(quote (1 1))", nil},
		{"'1", "(quote 1)", nil},
		{"'", "", anyError},
	})
}

func TestReadStringBackquote(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"`(1 1)", "(` (1 1))", nil},
		{"`1", "(` 1)", nil},
		{"`", "", anyError},
	})
}

func TestReadStringCharacter(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"?a", "97", nil},
		{"?9", "57", nil},
		{"?\\9", "57", nil},
		{"?\\999", "", anyError},
		{"?999", "", anyError},
		{"? ", "32", nil},
		{"?", "", anyError},
		{"?\\n", "10", nil},
		{"?\\\n", "-1", nil},
		{"?\\r", "13", nil},
		{"?\\ ", "32", nil},
		{"?\\71", "57", nil},
		{"?\\071", "57", nil},
		{"?\\0071", "", anyError},
	})
}

func TestReadStringComment(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"(+ 1 1) ; foo", "(+ 1 1)", nil},
		{"(+ 1 1) ;; foo", "(+ 1 1)", nil},
		{"(+ 1 1) ;; foo\n", "(+ 1 1)", nil},
		{"(+ 1 1) ;;; foo", "(+ 1 1)", nil},
		{"(+ 1 1);foo", "(+ 1 1)", nil},
		{"(+ 1 1);", "(+ 1 1)", nil},
		{";(+ 1 1)", "", anyError},
		{"  ;(+ 1 1)", "", anyError},
	})
}

func TestReadStringWhitespace(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{"  1", "1", nil},
		{"\t1", "1", nil},
		{"\n\n1", "1", nil},
		{"\t1  ", "1", nil},
		{"", "", anyError},
	})
}

func TestReadStringStr(t *testing.T) {
	testReadStringHelper(t, []readStringTestCase{
		{`""`, `""`, nil},
		{`".1"`, `".1"`, nil},
		{`"hello"`, `"hello"`, nil},
		{`"'"`, `"'"`, nil},
		{`"''"`, `"''"`, nil},
		{"\"hello\nworld\"", "\"hello\nworld\"", nil},
		{`"hello\041"`, `"hello!"`, nil},
		{"\"hello \\\nworld\"", `"hello world"`, nil},
		{`"say \"hello\""`, `"say "hello""`, nil},
		{`"hello`, "", anyError},
		{`"`, "", anyError},
		{"\"hello \\\n", "", anyError},
	})
}
