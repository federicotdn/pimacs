;;; goroutine_test.el --- Tests for goroutine.go  -*- lexical-binding: t; -*-

(lt--deftest test-reverse ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (defconst global nil)
    (lt--should-not global "no value")
    (defun mytest ()
      (setq global 42))
    (pimacs-go 'mytest)
    (pimacs-sleep 300)
    (lt--should (equal global 42) "equals 42")
    ))
