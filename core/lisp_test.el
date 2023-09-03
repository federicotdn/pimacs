;;; lisp_test.el --- Like 0.1% of ERT  -*- lexical-binding: t; -*-

;; Bootstrap some utility and testing functions
;; TODO: In the future, switch to normal defmacro/defun, or even
;; try loading ERT.

;; A very simplified version of defmacro
(defalias 'lt--defmacro
  (cons
   'macro
   #'(lambda (name arglist &rest body)
       "Define NAME as a macro."
       (list 'defalias
	     (list 'quote name)
	     (list 'cons ''macro (list 'function (cons 'lambda (cons arglist body))))))))

;; A very simplified version of defun
(lt--defmacro lt--defun (name arglist &rest body)
  (list 'defalias
	(list 'quote name)
	(list 'function (cons 'lambda (cons arglist body)))))

(setq lt--tests-to-run nil)
(setq lt--failure-info nil)
(setq lt--dbg-count 0)

(lt--defmacro lt--deftest (name arglist &rest body)
  (setq lt--tests-to-run (cons name lt--tests-to-run))
  (list 'defalias
	(list 'quote name)
	(list 'function (cons 'lambda (cons arglist body)))))

(lt--defun lt--should (val &optional msg)
  (if (null val)
    (signal 'assertion-failure (if msg msg "(no details)"))))

(lt--defun lt--should-not (val &optional msg)
  (lt--should (null val) msg))

(lt--defun lt--debug (val)
  (princ "!!! DEBUG " t)
  (princ lt--dbg-count t)
  (princ ": " t)
  (princ val t)
  (princ "\n" t)
  (setq lt--dbg-count (+ lt--dbg-count 1)))

(lt--defun lt--run-all-tests ()
  (while lt--tests-to-run
    (setq test (car lt--tests-to-run))
    (princ "+++ RUN-LISP: " t)
    (princ (symbol-name test) t)
    (princ "\n" t)
    (unwind-protect (funcall test)
      (setq lt--failure-info (symbol-name test)))
    (setq lt--tests-to-run (cdr lt--tests-to-run))))
