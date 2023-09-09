;;; goroutine_test.el --- Tests for goroutine.go  -*- lexical-binding: t; -*-

(lt--deftest test-pimacs-go-simple ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (defconst global nil)
    (lt--should-not global "no value")
    (defun mytest ()
      (setq global 42))
    (pimacs-go 'mytest)
    (pimacs-sleep 300)
    (lt--should (equal global 42) "equals 42")))

(lt--deftest test-pimacs-go-channel ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (setq ch (pimacs-chan 1))
    (defun mytest ()
      (pimacs-send ch 42))
    (pimacs-go 'mytest)
    (lt--should (equal (car (pimacs-receive ch)) 42)
		"equals 42")
    (pimacs-close ch)
    (lt--should-not (cdr (pimacs-receive ch)) "closed channel")))
