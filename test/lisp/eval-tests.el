;;; eval-tests.el --- Tests for eval.go  -*- lexical-binding: t; -*-

(lt--deftest test-let-closure ()
  (setq cl nil)
  (let ((a "hello"))
    (fset 'cl (function (lambda () (cons a a)))))

  (lt--should (equal (cl) (cons "hello" "hello"))
	      "call cl gives (hello . hello)"))

(lt--deftest test-let* ()
  (let* (foo
	 (a 1)
	 (b (+ a 1))
	 (c (+ b 10))
	 (d "hello"))
    (lt--should-not foo "foo is nil")
    (lt--should (equal a 1) "a is 1")
    (lt--should (equal b 2) "b is 2")
    (lt--should (equal c 12) "b is 12")
    (lt--should (equal d "hello") "d is hello"))

  (lt--should-not (boundp 'a) "a is not bound"))

(lt--deftest test-let*-closure ()
  (setq cl nil)
  (let* ((a 1)
	 (b (+ a 1))
	 (c (+ b 10)))
    (fset 'cl (function (lambda () (cons b c)))))

  (lt--should (equal (cl) (cons 2 12))
	      "call cl gives (2 . 12)"))
