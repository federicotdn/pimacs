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

;; A very simplified version of defconst
(defmacro defconst (name value &optional _docs)
  (list 'setq name value))

;;; pimacs.el ends here
