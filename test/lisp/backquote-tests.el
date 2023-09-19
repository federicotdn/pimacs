;;; backquote-tests.el --- Tests for Emacs backquote.el  -*- lexical-binding: t; -*-

;; Tests adapted from Emacs' backquote-tests.el
;; TODO: Replace for original test file

(lt--deftest backquote-test-basic ()
  (let ((lst '(ba bb bc))
        (vec [ba bb bc]))
    (lt--should (equal 3 `,(eval '(+ x y) '((x . 1) (y . 2)))) "t1")
    ;; (lt--should (equal vec `[,@lst]) "t2")
    (lt--should (equal `(a lst c) '(a lst c)) "t3")
    (lt--should (equal `(a ,lst c) '(a (ba bb bc) c)) "t4")
    ;; (lt--should (equal `(a ,@lst c) '(a ba bb bc c)) "t5")
    (lt--should (equal `(a vec c) '(a vec c)) "t6")
    (lt--should (equal `(a ,vec c) '(a [ba bb bc] c)) "t7")
    ;; (lt--should (equal `(a ,@vec c) '(a ba bb bc c)) "t8")
    ))

(lt--deftest backquote-test-nested ()
  "Test nested backquotes."
  (let ((lst '(ba bb bc))
        (vec [ba bb bc]))
    ;; (lt--should (equal `(a ,`(,@lst) c) `(a ,lst c)) "t1")
    ;; (lt--should (equal `(a ,`[,@lst] c) `(a ,vec c)) "t2")
    ;; (lt--should (equal `(a ,@`[,@lst] c) `(a ,@lst c)) "t3")
    ))
