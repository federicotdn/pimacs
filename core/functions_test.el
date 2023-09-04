;;; functions_test.el --- Tests for functions.go  -*- lexical-binding: t; -*-

(lt--deftest test-reverse ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (lt--should (equal (reverse '(1 2 3 4))
		       '(4 3 2 1))
		"list reversed")
    (lt--should (equal (reverse ()) ()) "empty list reversed")
    (lt--should (equal (reverse [1 2 3 4])
		       [4 3 2 1])
		"vec reversed")
    (lt--should (equal (reverse []) [])
		"empty vec reversed")))
