;;; lt.el --- Like 0.1% of ERT  -*- lexical-binding: t; -*-

;; lt == Lisp Testing
;; TODO: In the future, switch to ERT if possible.

(setq lt--tests-to-run nil)
(setq lt--failure-info nil)
(setq lt--dbg-count 0)

(defmacro lt--deftest (name arglist &rest body)
  (declare (indent 2))
  (if (memq name lt--tests-to-run)
      (signal 'test-duplicated (symbol-name name)))
  (setq lt--tests-to-run (cons name lt--tests-to-run))
  (list 'defalias
	(list 'quote name)
	(list 'function (cons 'lambda (cons arglist body)))))

(defun lt--should (val &optional msg)
  (if (null val)
    (signal 'assertion-failure (if msg msg "(no details)"))))

(defun lt--should-not (val &optional msg)
  (lt--should (null val) msg))

(defun lt--debug (val)
  (princ "!!! DEBUG " t)
  (princ lt--dbg-count t)
  (princ ": " t)
  (princ val t)
  (princ "\n" t)
  (setq lt--dbg-count (+ lt--dbg-count 1)))

(defun lt--run-all-tests ()
  (setq lt--tests-to-run (reverse lt--tests-to-run))
  (while lt--tests-to-run
    (setq test (car lt--tests-to-run))
    (princ "+++ LISP: " t)
    (princ (symbol-name test) t)
    (princ "\n" t)
    (unwind-protect (funcall test)
      (setq lt--failure-info (symbol-name test)))
    (setq lt--tests-to-run (cdr lt--tests-to-run))))

;;; lt.el ends here
