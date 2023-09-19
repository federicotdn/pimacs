;;; loadup-pimacs.el --- Base file for Pimacs  -*- lexical-binding: t; -*-

;; Eventually to be replaced with loadup.el

(load "emacs-lisp/debug-early.el")
(load "emacs-lisp/byte-run.el")
(load "emacs-lisp/backquote.el")

;; TODO: Remove once subr.el is loaded
(defmacro unless (cond &rest body)
  (cons 'if (cons cond (cons nil body))))
(defalias 'not #'null)
(defmacro push (newelt place)
  (if (symbolp place)
      (list 'setq place
            (list 'cons newelt place))
    (require 'macroexp)
    (macroexp-let2 macroexp-copyable-p x newelt
      (gv-letplace (getter setter) place
        (funcall setter `(cons ,x ,getter))))))

;; (load "subr.el")

;;; pimacs.el ends here
