;;; pimacs.el --- Base file for Pimacs  -*- lexical-binding: t; -*-

;; A very simplified version of defmacro
(defalias 'pimacs-defmacro
  (cons
   'macro
   #'(lambda (name arglist &rest body)
       "Define NAME as a macro."
       (list 'defalias
	     (list 'quote name)
	     (list 'cons ''macro (list 'function (cons 'lambda (cons arglist body))))))))

;; A very simplified version of defun
(pimacs-defmacro pimacs-defun (name arglist &rest body)
  (list 'defalias
	(list 'quote name)
	(list 'function (cons 'lambda (cons arglist body)))))

;; Set up some Pimacs-related variables/functions
(setq pimacs-testing (getenv-internal "PIMACS_TESTING"))

(pimacs-defun pimacs-greet ()
  (let ((msg "Welcome to Pimacs!\n"))
    (princ msg t)
    nil))

(if (null pimacs-testing)
    (pimacs-greet))

;;; pimacs.el ends here
