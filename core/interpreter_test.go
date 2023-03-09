package core

import (
	"fmt"
	"testing"
)

type stringToStringTC struct {
	input    string
	expected string
	err      error
}

var anyError = fmt.Errorf("anyError")

func testStringToString(t *testing.T, fn func(string) (string, error), cases []stringToStringTC) {
	for _, tc := range cases {
		if tc.err != nil && tc.expected != "" {
			t.Fatalf("input: '%v' defines both expected value and error", tc.input)
		}

		output, err := fn(tc.input)
		if err != nil {
			if tc.err == nil {
				t.Logf("input: '%v' got error '%v'", tc.input, err)
				t.Fail()
			} else if tc.err != anyError && tc.err != err {
				t.Logf("input: '%v' got error '%v' instead of error '%v'", tc.input, err, tc.err)
				t.Fail()
			}
		} else {
			if tc.err != nil {
				t.Logf("input: '%v' expected error '%v' instead of value '%v'", tc.input, tc.err, output)
				t.Fail()
			} else if output != tc.expected {
				t.Logf("input: '%v' expected '%v' but got '%v'", tc.input, tc.expected, output)
				t.Fail()
			}
		}
	}
}

func readPrin1(input string) (string, error) {
	inp := NewInterpreter()
	return inp.ReadPrin1(input)
}

func readEvalPrin1(input string) (string, error) {
	inp := NewInterpreter()
	return inp.ReadEvalPrin1(input)
}

func TestReadPrint(t *testing.T) {
	testStringToString(t, readPrin1, []stringToStringTC{
		{"1", "1", nil},
		{"-1", "-1", nil},
		{"0", "0", nil},
		{"-0", "0", nil},
		{"1.2", "1.2", nil},
		{"0.111", "0.111", nil},
		{".11", "0.11", nil},
		{".hello", ".hello", nil},
		{"123.123e-2", "1.23123", nil},
		{"123.123E-2", "1.23123", nil},
		{"nil", "nil", nil},
		{"()", "nil", nil},
		{"(1 . 1)", "(1 . 1)", nil},
		{"( . 2)", "2", nil},
		{"(. 2)", "2", nil},
		{"(. .test)", ".test", nil},
		{"(.test .)", "(.test .)", nil},
		{"(. .)", ".", nil},
		{"( . .)", ".", nil},
		{"(1 . (2 . 3))", "(1 2 . 3)", nil},
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
		{"foo", "foo", nil},
		{"foo\\ bar", "foo bar", nil},
		{"foo:bar", "foo:bar", nil},
		{"foo-bar", "foo-bar", nil},
		{"foo-bar?", "foo-bar?", nil},
		{"\\99", "99", nil},
		{"\\", "", anyError},
		{"foo\\", "", anyError},
		{"'(1 1)", "(quote (1 1))", nil},
		{"'1", "(quote 1)", nil},
		{"'", "", anyError},
		{"`(1 1)", "(` (1 1))", nil},
		{"`1", "(` 1)", nil},
		{"`", "", anyError},
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
		{"(+ 1 1) ; foo", "(+ 1 1)", nil},
		{"(+ 1 1) ;; foo", "(+ 1 1)", nil},
		{"(+ 1 1) ;; foo\n", "(+ 1 1)", nil},
		{"(+ 1 1) ;;; foo", "(+ 1 1)", nil},
		{"(+ 1 1);foo", "(+ 1 1)", nil},
		{"(+ 1 1);", "(+ 1 1)", nil},
		{";(+ 1 1)", "", anyError},
		{"  ;(+ 1 1)", "", anyError},
		{"  1", "1", nil},
		{"\t1", "1", nil},
		{"\n\n1", "1", nil},
		{"\t1  ", "1", nil},
		{"", "", anyError},
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

func TestReadEvalPrint(t *testing.T) {
	testStringToString(t, readEvalPrin1, []stringToStringTC{
		{"1", "1", nil},
		{"\"hello\"", "\"hello\"", nil},
		{"\"ñandú\"", "\"ñandú\"", nil},
		{"nil", "nil", nil},
		{"'nil", "nil", nil},
		{"t", "t", nil},
		{"'t", "t", nil},
		{"'-", "-", nil},
		{"+", "", anyError},
		{"(+)", "0", nil},
		{"(+ 1)", "1", nil},
		{"(+ 1 1)", "2", nil},
		{"(+ 1 (+ 1 1))", "3", nil},
		{"(+ 1 nil)", "", anyError},
		{"(memq 'a '(a b c))", "(a b c)", nil},
		{"(memq 'a '(a . b))", "(a . b)", nil},
		{"(memq 'c '(a . b))", "", anyError},
		{"(length '(1 2 3))", "3", nil},
		{"(length '(1 2 . 3))", "", anyError},
		{"(eval '(+ 1 1))", "2", nil},
		{"(progn 1 2 3)", "3", nil},
		{"(progn)", "nil", nil},
		{"(progn (set 'a 42) a)", "42", nil},
		{"(length \"hello\")", "5", nil},
		{"foo", "", anyError},
		{"(foo 1)", "", anyError},
		{"(list 1 2 3)", "(1 2 3)", nil},
		{"(list)", "nil", nil},
		{"(cons)", "", anyError},
		{"(cons 1 2 3)", "", anyError},
		{"(if t 1 2)", "1", nil},
		{"(if nil 1 2)", "2", nil},
		{"(if t 42)", "42", nil},
		{"(if nil 42)", "nil", nil},
		{"unbound", "", anyError},
		{"(progn (set 'f '((b . 42))) (eval 'b f))", "42", nil},
		{"(eval)", "", anyError},
		{"(eval '(+ 1 1))", "2", nil},
		{"(eq nil nil)", "t", nil},
		{"(eq 'nil nil)", "t", nil},
		{"(eq nil ())", "t", nil},
		{"(eq t t)", "t", nil},
		{"(eq 't t)", "t", nil},
		{"(eq 1000 1000)", "nil", nil},
		{"(consp nil)", "nil", nil},
		{"(consp (cons 1 2))", "t", nil},
		{"(put 1 1)", "", anyError},
		{"(put 'a 'foo 1)", "1", nil},
		{"(progn (put 'a 'foo 42) (get 'a 'foo))", "42", nil},
		{"(progn (if nil 1 2 3 (set 'foo 42)) foo)", "42", nil},
		{"(throw 'f 1)", "", anyError},
		{"(catch 'f (throw 'f 1))", "1", nil},
		{"(catch 'k (throw 'f 1))", "", anyError},
		{"(catch)", "", anyError},
		{"(catch 'k)", "nil", nil},
		{"(progn (catch 'a (unwind-protect (throw 'a nil) (set 'foo 42))) foo)", "42", nil},
		{"(symbol-plist 'error)", "(error-conditions (error) error-message \"error\")", nil},
		{"(symbol-plist 'quit)", "(error-conditions (quit) error-message \"Quit\")", nil},
		{"(symbol-plist 'user-error)", "(error-conditions (user-error error) error-message \"\")", nil},
		{"(plist-put '(a 1 b 2) 'a 3)", "(a 3 b 2)", nil},
		{"(plist-put '(a 1 b 2) 'c 3)", "(a 1 b 2 c 3)", nil},
		{"(plist-put nil 'c 3)", "(c 3)", nil},
		{"(condition-case nil 42)", "42", nil},
		{"(condition-case foo 42)", "42", nil},
		{"(condition-case nil (signal 'user-error nil))", "", anyError},
		{"(condition-case nil (signal 'user-error nil) (user-error 42))", "42", nil},
		{"(condition-case nil (signal 'user-error nil) (user-error 1) (user-error 2))", "1", nil},
		{"(condition-case nil (throw 'foo 1) (t 42))", "42", nil},
		{"(catch 'foo (condition-case nil (throw 'foo 1234) (t 42)))", "1234", nil},
		{"(condition-case a (signal 'error \"foo\") ((error) a))", "(error . \"foo\")", nil},
		{"(<)", "", anyError},
		{"(< 1 4)", "t", nil},
		{"(< 4 2)", "nil", nil},
		{"(< 4 4)", "nil", nil},
		{"(symbol-name '##)", `""`, nil},
		{"(equal 1 1)", "t", nil},
		{"(equal 1 2)", "nil", nil},
		{"(equal '(1 2 3) '(1 2 3))", "t", nil},
		{"(progn (fset 'foo (function (lambda (x) (+ x 1)))) (foo 1))", "2", nil},
		{"(progn (fset 'foo (function (lambda (x) (+ x 1)))) (funcall 'foo 1))", "2", nil},
		{"(progn (fset 'foo (function (lambda () \"foo\"))) (foo 1))", "", anyError},
		{"(progn (fset 'foo (function (lambda () \"foo\"))) (foo))", `"foo"`, nil},
		{"(progn (fset 'foo (function (lambda (&optional x) x))) (foo))", "nil", nil},
		{"(progn (fset 'foo (function (lambda (&rest x) x))) (foo 1 2 3))", "(1 2 3)", nil},
		{"(funcall '+)", "0", nil},
		{"(funcall '+ 1 2 3)", "6", nil},
		{"(funcall '+ 1 2 (+ 1 1))", "5", nil},
		{"(funcall 'list 1 2)", "(1 2)", nil},
		{"(funcall 'if t 1)", "", anyError},
		{"(apply 'cons)", "", anyError},
		{"(apply 'cons '(1 2))", "(1 . 2)", nil},
		{"(apply 'cons '(1))", "", anyError},
		{"(apply '+ 1 '(2 3 4))", "10", nil},
		{"(apply 'if '(t 1)", "", anyError},
		{"(apply 'eval '(t))", "t", nil},
		{`(progn (insert "foo") (buffer-string))`, `"foo"`, nil},
		{`(buffer-name (current-buffer))`, `"*scratch*"`, nil},
		{`(assq 'a '((a . b) (c . d)))`, `(a . b)`, nil},
		{`(assq 'f '((a . b) (c . d)))`, `nil`, nil},
		{`(assq 'f '((a . b) (c . d) . 123))`, "", anyError},
		{`(nconc)`, `nil`, nil},
		{`(nconc 1)`, `1`, nil},
		{`(nconc nil)`, `nil`, nil},
		{`(nconc '(1 2 3) '(4 5 6))`, `(1 2 3 4 5 6)`, nil},
		{`(nconc '(1 2 3) nil '(4 5 6))`, `(1 2 3 4 5 6)`, nil},
		{`(nconc '(1 2 3) '(4 5 6) 7)`, `(1 2 3 4 5 6 . 7)`, nil},
		{`(nconc (cons 1 2) '(4 5 6))`, `(1 4 5 6)`, nil},
		{`(length (buffer-list))`, `1`, nil},
		{"(getenv-internal \"FOO\")", "nil", nil},
		{"(getenv-internal \"PIMACS_TESTING\")", "\"true\"", nil},
	})
}
