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

(lt--deftest test-eq ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (lt--should (eq 'a 'a) "a eq a")
    (lt--should-not (eq 'a 'b) "not a eq b")
    (lt--should-not (eq 1 1) "not 1 eq 1")
    (let ((foo "hello"))
      (lt--should (eq foo foo) "foo eq foo"))))

(lt--deftest test-eql ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (lt--should (eql 'a 'a) "a eql a")
    (lt--should (eql 1 1) "1 eql 1")
    (lt--should-not (eql "hello" "hello") "not hello eql hello")))

(lt--deftest test-equal ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (lt--should (equal 'a 'a) "a equal a")
    (lt--should (equal 1 1) "1 equal 1")
    (lt--should (equal "hello" "hello") "hello equal hello")))
