package core

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type stringToStringTC struct {
	input    string
	expected string
	err      error
	inp      *Interpreter
}

var anyError = fmt.Errorf("anyError")
var globalInp = newTestingInterpreter()

func testStringToString(t *testing.T, fn func(string, *Interpreter) (string, error), cases []stringToStringTC) {
	for _, tc := range cases {
		if tc.err != nil && tc.expected != "" {
			t.Fatalf("input: '%v' defines both expected value and error", tc.input)
		}

		output, err := fn(tc.input, tc.inp)
		if err != nil {
			if tc.err == nil {
				t.Logf("input: '%v' got error '%v'", tc.input, err)
				t.Fail()
			} else if !errors.Is(tc.err, anyError) && !errors.Is(tc.err, err) {
				t.Logf(
					"input: '%v' got error '%v' instead of error '%v'",
					tc.input,
					strings.Split(err.Error(), "\n")[0],
					strings.Split(tc.err.Error(), "\n")[0],
				)
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

func readPrin1(input string, inp *Interpreter) (string, error) {
	return globalInp.ReadPrin1(input)
}

func readEvalPrin1(input string, inp *Interpreter) (string, error) {
	if inp == nil {
		inp = globalInp
	}
	return inp.ReadEvalPrin1(input)
}

func TestReadPrint(t *testing.T) {
	testStringToString(t, readPrin1, []stringToStringTC{
		{"1", "1", nil, nil},
		{"-1", "-1", nil, nil},
		{"0", "0", nil, nil},
		{"-0", "0", nil, nil},
		{"1.2", "1.2", nil, nil},
		{"0.111", "0.111", nil, nil},
		{".11", "0.11", nil, nil},
		{".hello", "\\.hello", nil, nil},
		{"123.123e-2", "1.23123", nil, nil},
		{"123.123E-2", "1.23123", nil, nil},
		{"nil", "nil", nil, nil},
		{"()", "nil", nil, nil},
		{"(1 . 1)", "(1 . 1)", nil, nil},
		{"( . 2)", "", anyError, nil},
		{"(. 2)", "", anyError, nil},
		{"(1 2 3 . 4)", "(1 2 3 . 4)", nil, nil},
		{"(. .test)", "", anyError, nil},
		{"(.test .)", "(\\.test \\.)", nil, nil},
		{"(. .)", "", anyError, nil},
		{"( . .)", "", anyError, nil},
		{"(1 . (2 . 3))", "(1 2 . 3)", nil, nil},
		{"(1 1)", "(1 1)", nil, nil},
		{"(1 2 .)", "(1 2 \\.)", nil, nil},
		{"( 1 1 )", "(1 1)", nil, nil},
		{"(.1 2)", "(0.1 2)", nil, nil},
		{"(.)", "(\\.)", nil, nil},
		{"((1 1))", "((1 1))", nil, nil},
		{"((1 2 3) (4 5 6))", "((1 2 3) (4 5 6))", nil, nil},
		{"(1 . (2 . (3 . nil)))", "(1 2 3)", nil, nil},
		{"(", "", anyError, nil},
		{"(123", "", anyError, nil},
		{"(123 1 2 3", "", anyError, nil},
		{"(123 1 2 3(", "", anyError, nil},
		{"((())", "", anyError, nil},
		{")", "", anyError, nil},
		{"())", "nil", nil, nil},
		{"foo", "foo", nil, nil},
		{"foo\\ bar", "foo bar", nil, nil},
		{"foo:bar", "foo:bar", nil, nil},
		{"foo-bar", "foo-bar", nil, nil},
		{"foo-bar?", "foo-bar?", nil, nil},
		{"\\99", "\\99", nil, nil},
		{"\\-10", "\\-10", nil, nil},
		{"\\+10", "\\+10", nil, nil},
		{"\\1.2", "\\1.2", nil, nil},
		{"\\+", "+", nil, nil},
		{"\\", "", anyError, nil},
		{"foo\\", "", anyError, nil},
		{"'(1 1)", "(quote (1 1))", nil, nil},
		{"'1", "(quote 1)", nil, nil},
		{"'", "", anyError, nil},
		{"`(1 1)", "(` (1 1))", nil, nil},
		{"`1", "(` 1)", nil, nil},
		{"`", "", anyError, nil},
		{"?a", "97", nil, nil},
		{"?9", "57", nil, nil},
		{"?\\9", "57", nil, nil},
		{"?\\999", "", anyError, nil},
		{"?999", "", anyError, nil},
		{"? ", "32", nil, nil},
		{"?", "", anyError, nil},
		{"?\\n", "10", nil, nil},
		{"?\\\n", "", anyError, nil},
		{"?\\r", "13", nil, nil},
		{"?\\ ", "32", nil, nil},
		{"?\\71", "57", nil, nil},
		{"?\\071", "57", nil, nil},
		{"?\\0071", "", anyError, nil},
		{"(+ 1 1) ; foo", "(+ 1 1)", nil, nil},
		{"(+ 1 1) ;; foo", "(+ 1 1)", nil, nil},
		{"(+ 1 1) ;; foo\n", "(+ 1 1)", nil, nil},
		{"(+ 1 1) ;;; foo", "(+ 1 1)", nil, nil},
		{"(+ 1 1);foo", "(+ 1 1)", nil, nil},
		{"(+ 1 1);", "(+ 1 1)", nil, nil},
		{";(+ 1 1)", "", anyError, nil},
		{"  ;(+ 1 1)", "", anyError, nil},
		{"  1", "1", nil, nil},
		{"\t1", "1", nil, nil},
		{"\n\n1", "1", nil, nil},
		{"\t1  ", "1", nil, nil},
		{"", "", anyError, nil},
		{`""`, `""`, nil, nil},
		{`".1"`, `".1"`, nil, nil},
		{`"hello"`, `"hello"`, nil, nil},
		{`"'"`, `"'"`, nil, nil},
		{`"''"`, `"''"`, nil, nil},
		{"\"hello\nworld\"", "\"hello\nworld\"", nil, nil},
		{`"hello\041"`, `"hello!"`, nil, nil},
		{"\"hello \\\nworld\"", `"hello world"`, nil, nil},
		{`"say \"hello\""`, `"say "hello""`, nil, nil},
		{`"hello`, "", anyError, nil},
		{`"`, "", anyError, nil},
		{"\"hello \\\n", "", anyError, nil},
		{"[]", "[]", nil, nil},
		{"[1]", "[1]", nil, nil},
		{"[1 2]", "[1 2]", nil, nil},
		{"[1 2 ]", "[1 2]", nil, nil},
		{"[     1 2 ]", "[1 2]", nil, nil},
		{"[1 2 (+ 1 1)]", "[1 2 (+ 1 1)]", nil, nil},
		{"[", "", anyError, nil},
		{"]", "", anyError, nil},
	})
}

func TestReadEvalPrint(t *testing.T) {
	testStringToString(t, readEvalPrin1, []stringToStringTC{
		{"1", "1", nil, nil},
		{"\"hello\"", "\"hello\"", nil, nil},
		{"\"ñandú\"", "\"ñandú\"", nil, nil},
		{"nil", "nil", nil, nil},
		{"'nil", "nil", nil, nil},
		{"t", "t", nil, nil},
		{"'t", "t", nil, nil},
		{"'-", "-", nil, nil},
		{"'(1 . 2)", "(1 . 2)", nil, nil},
		{"+", "", anyError, nil},
		{"(+)", "0", nil, nil},
		{"(*)", "1", nil, nil},
		{"(* 1 a)", "", anyError, nil},
		{"(* 0)", "0", nil, nil},
		{"(* 3)", "3", nil, nil},
		{"(* -1 2 -3)", "6", nil, nil},
		{"(* 3.14 2 -3)", "", anyError, nil}, // NOTE: this will fail once arithmetic operators are implemented for floats
		{"(% 10 2)", "0", nil, nil},
		{"(% 10 (+ 10 3))", "10", nil, nil},
		{"(% 12332132122 10)", "2", nil, nil},
		{"(eval '(% 2 10))", "2", nil, nil},
		{"(% 12342 -10)", "2", nil, nil},
		{"(% 2 a)", "", anyError, nil},
		{"(% 2 1.2)", "", anyError, nil},
		{"(% 2)", "", anyError, nil},
		{"(+ 1)", "1", nil, nil},
		{"(+ 1 1)", "2", nil, nil},
		{"(+ 1 (+ 1 1))", "3", nil, nil},
		{"(+ 1 nil)", "", anyError, nil},
		{"(memq 'a '(a b c))", "(a b c)", nil, nil},
		{"(memq 'a '(a . b))", "(a . b)", nil, nil},
		{"(memq 'c '(a . b))", "", anyError, nil},
		{"(length '(1 2 3))", "3", nil, nil},
		{"(length '(1 2 . 3))", "", anyError, nil},
		{"(eval '(+ 1 1))", "2", nil, nil},
		{"(progn 1 2 3)", "3", nil, nil},
		{"(progn)", "nil", nil, nil},
		{"(progn (set 'a 42) a)", "42", nil, nil},
		{"(length \"hello\")", "5", nil, nil},
		{"(length \"helláá\")", "6", nil, nil},
		{"foo", "", anyError, nil},
		{"(foo 1)", "", anyError, nil},
		{"(list 1 2 3)", "(1 2 3)", nil, nil},
		{"(list)", "nil", nil, nil},
		{"(cons)", "", anyError, nil},
		{"(cons 1 2 3)", "", anyError, nil},
		{"(if t 1 2)", "1", nil, nil},
		{"(if nil 1 2)", "2", nil, nil},
		{"(if t 42)", "42", nil, nil},
		{"(if nil 42)", "nil", nil, nil},
		{"unbound", "", anyError, nil},
		{"(boundp 'asdf)", "nil", nil, nil},
		{"(progn (setq aaa 1) (boundp 'aaa))", "t", nil, nil},
		{"(fboundp 'asdf)", "nil", nil, nil},
		{"(fboundp 'length)", "t", nil, nil},
		{"(progn (set 'f '((b . 42))) (eval 'b f))", "42", nil, nil},
		{"(eval)", "", anyError, nil},
		{"(eval '(+ 1 1))", "2", nil, nil},
		{"(eq nil nil)", "t", nil, nil},
		{"(eq 'nil nil)", "t", nil, nil},
		{"(eq nil ())", "t", nil, nil},
		{"(eq t t)", "t", nil, nil},
		{"(eq 't t)", "t", nil, nil},
		{"(eq 1000 1000)", "nil", nil, nil},
		{"(consp nil)", "nil", nil, nil},
		{"(consp (cons 1 2))", "t", nil, nil},
		{"(put 1 1)", "", anyError, nil},
		{"(put 'a 'foo 1)", "1", nil, nil},
		{"(progn (put 'a 'foo 42) (get 'a 'foo))", "42", nil, nil},
		{"(progn (if nil 1 2 3 (set 'foo 42)) foo)", "42", nil, nil},
		{"(throw 'f 1)", "", anyError, nil},
		{"(catch 'f (throw 'f 1))", "1", nil, nil},
		{"(catch 'k (throw 'f 1))", "", anyError, nil},
		{"(catch)", "", anyError, nil},
		{"(catch 'k)", "nil", nil, nil},
		{"(progn (catch 'a (unwind-protect (throw 'a nil) (set 'foo 42))) foo)", "42", nil, nil},
		{"(symbol-plist 'error)", "(error-conditions (error) error-message \"error\")", nil, nil},
		{"(symbol-plist 'quit)", "(error-conditions (quit) error-message \"Quit\")", nil, nil},
		{"(symbol-plist 'user-error)", "(error-conditions (user-error error) error-message \"\")", nil, nil},
		{"(plist-put '(a 1 b 2) 'a 3)", "(a 3 b 2)", nil, nil},
		{"(plist-put '(a 1 b 2) 'c 3)", "(a 1 b 2 c 3)", nil, nil},
		{"(plist-put nil 'c 3)", "(c 3)", nil, nil},
		{"(plist-put nil :foo 1)", "(:foo 1)", nil, nil},
		{"(plist-put (cons 1 2) :a 1)", "", anyError, nil},
		{"(plist-put 99 :a 1)", "", anyError, nil},
		{"(condition-case nil 42)", "42", nil, nil},
		{"(condition-case foo 42)", "42", nil, nil},
		{"(condition-case nil (signal 'user-error nil))", "", anyError, nil},
		{"(condition-case nil (signal 'user-error nil) (user-error 42))", "42", nil, nil},
		{"(condition-case nil (signal 'user-error nil) (user-error 1) (user-error 2))", "1", nil, nil},
		{"(condition-case nil (throw 'foo 1) (t 42))", "42", nil, nil},
		{"(catch 'foo (condition-case nil (throw 'foo 1234) (t 42)))", "1234", nil, nil},
		{"(condition-case a (signal 'error \"foo\") ((error) a))", "(error . \"foo\")", nil, nil},
		{"(<)", "", anyError, nil},
		{"(< 1 4)", "t", nil, nil},
		{"(< 4 2)", "nil", nil, nil},
		{"(< 4 4)", "nil", nil, nil},
		{"(symbol-name '##)", `""`, nil, nil},
		{"(equal 1 1)", "t", nil, nil},
		{"(equal 1 2)", "nil", nil, nil},
		{"(equal '(1 2 3) '(1 2 3))", "t", nil, nil},
		{"(progn (fset 'foo (function (lambda (x) (+ x 1)))) (foo 1))", "2", nil, nil},
		{"(progn (fset 'foo (function (lambda (x) (+ x 1)))) (funcall 'foo 1))", "2", nil, nil},
		{"(progn (fset 'foo (function (lambda () \"foo\"))) (foo 1))", "", anyError, nil},
		{"(progn (fset 'foo (function (lambda () \"foo\"))) (foo))", `"foo"`, nil, nil},
		{"(progn (fset 'foo (function (lambda (&optional x) x))) (foo))", "nil", nil, nil},
		{"(progn (fset 'foo (function (lambda (&optional x y) x))) (foo))", "nil", nil, nil},
		{"(progn (fset 'foo (function (lambda (&rest x) x))) (foo 1 2 3))", "(1 2 3)", nil, nil},
		{"(funcall '+)", "0", nil, nil},
		{"(funcall '+ 1 2 3)", "6", nil, nil},
		{"(funcall '+ 1 2 (+ 1 1))", "5", nil, nil},
		{"(funcall 'list 1 2)", "(1 2)", nil, nil},
		{"(funcall 'if t 1)", "", anyError, nil},
		{"(apply 'cons)", "", anyError, nil},
		{"(apply 'cons '(1 2))", "(1 . 2)", nil, nil},
		{"(apply 'cons '(1))", "", anyError, nil},
		{"(apply '+ 1 '(2 3 4))", "10", nil, nil},
		{"(apply 'if '(t 1)", "", anyError, nil},
		{"(apply 'eval '(t))", "t", nil, nil},
		{`(progn (insert "foo") (buffer-string))`, `"foo"`, nil, nil},
		{`(buffer-name (current-buffer))`, `"*scratch*"`, nil, nil},
		{`(assq 'a '((a . b) (c . d)))`, `(a . b)`, nil, nil},
		{`(assq 'f '((a . b) (c . d)))`, `nil`, nil, nil},
		{`(assq 'f '((a . b) (c . d) . 123))`, "", anyError, nil},
		{`(nconc)`, `nil`, nil, nil},
		{`(nconc 1)`, `1`, nil, nil},
		{`(nconc nil)`, `nil`, nil, nil},
		{`(nconc '(1 2 3) '(4 5 6))`, `(1 2 3 4 5 6)`, nil, nil},
		{`(nconc '(1 2 3) nil '(4 5 6))`, `(1 2 3 4 5 6)`, nil, nil},
		{`(nconc '(1 2 3) '(4 5 6) 7)`, `(1 2 3 4 5 6 . 7)`, nil, nil},
		{`(nconc (cons 1 2) '(4 5 6))`, `(1 4 5 6)`, nil, nil},
		{`(length (buffer-list))`, `1`, nil, nil},
		{"(getenv-internal \"FOO\")", "nil", nil, nil},
		{"(intern \"\")", "##", nil, nil},
		{`(intern "\\")`, `\\`, nil, nil},
		{"(setq)", "nil", nil, nil},
		{"(setq a 1)", "1", nil, nil},
		{"(setq a 1 b 2)", "2", nil, nil},
		{"(setq a 1 b 2 c)", "", anyError, nil},
		{"(progn (setq a 1) 1)", "1", nil, nil},
		{"(let)", "", anyError, nil},
		{"(let (a) a)", "nil", nil, nil},
		{"(let ((a nil)) a)", "nil", nil, nil},
		{"(let ((a nil) (b 2)) b)", "2", nil, nil},
		{"(let ((a nil) (b 2) c) c)", "nil", nil, nil},
		{"(progn (let (abc321)) abc321)", "", anyError, nil},
		{"noninteractive", "t", nil, nil},
		{"(cons (let ((noninteractive nil)) noninteractive) noninteractive)", "(nil . t)", nil, nil},
		{`(progn (setq pimacs--repo "hello") pimacs--repo)`, `"hello"`, nil, nil},
		{"(plist-get (pimacs--symbol-debug 'pimacs--repo) :special)", "t", nil, nil},
		{"[1 2 3]", "[1 2 3]", nil, nil},
		{"[[[]]]", "[[[]]]", nil, nil},
		{"[1 2 (+ 1 1)]", "[1 2 (+ 1 1)]", nil, nil},
		{"(vectorp [])", "t", nil, nil},
		{"(make-char-table nil)", "#^[nil nil 0 nil]", nil, nil},
	})
}

func TestReadEvalPrintSpecificErr(t *testing.T) {
	t.Parallel()
	inp := newTestingInterpreter()
	ec := inp.ec
	sentinel := newString("sentinel", false)

	// Note: all these will be run with the interpreter `inp`, not the global one
	cases := []stringToStringTC{
		{"(1 . 2)", "", xErrOnly(ec.wrongTypeArgument(ec.s.listp, sentinel)), inp},
		{"(foo 1 2 3)", "", xErrOnly(ec.signalN(ec.s.voidFunction)), inp},
		{")", "", xErrOnly(ec.signalN(ec.s.invalidReadSyntax)), inp},
		{"(char-table-range 123 123)", "", xErrOnly(ec.wrongTypeArgument(ec.s.charTablep, sentinel)), inp},
		{"(signal 1 2)", "", xErrOnly(ec.wrongTypeArgument(ec.s.symbolp, sentinel)), inp},
		{"(plist-put '(:b) :a 1)", "", xErrOnly(ec.wrongTypeArgument(ec.s.plistp, sentinel)), inp},
		{"(plist-put '(1 2 3 . 4) :a 1)", "", xErrOnly(ec.wrongTypeArgument(ec.s.plistp, sentinel)), inp},
		{"(make-hash-table :test)", "", xErrOnly(ec.wrongTypeArgument(ec.s.plistp, sentinel)), inp},
		{"(make-hash-table :test 'eq :foo)", "", xErrOnly(ec.wrongTypeArgument(ec.s.plistp, sentinel)), inp},
		{"(make-hash-table :test 'eq :foo 123)", "", xErrOnly(ec.signalError("")), inp},
		{"(make-hash-table :test 'foo)", "", xErrOnly(ec.signalError("")), inp},
	}

	testStringToString(t, readEvalPrin1, cases)
}
