;;; pimacs.el --- Base file for Pimacs  -*- lexical-binding: t; -*-

;; A very simplified version of defmacro
(defalias 'defmacro
  (cons
   'macro
   #'(lambda (name arglist &rest body)
       "Define NAME as a macro."
       (list 'defalias
	     (list 'quote name)
	     (list 'cons ''macro (list 'function (cons 'lambda (cons arglist body))))))))

;; A very simplified version of defun
(defmacro defun (name arglist &rest body)
  (list 'defalias
	(list 'quote name)
	(list 'function (cons 'lambda (cons arglist body)))))

;; A very simplified version of defvar
(defmacro defvar (name &optional value _docs)
  (list 'setq name value))

(load "emacs-lisp/debug-early.el")
;; (load "emacs-lisp/byte-run.el")
;; (load "emacs-lisp/backquote.el")
;; (load "subr.el")

;;; pimacs.el ends here
