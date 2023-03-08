;; Base file for Pimacs

(set 'pimacs-testing (getenv-internal "PIMACS_TESTING"))

(fset 'pimacs-greet-fn (function (lambda () (progn (princ "Welcome to Pimacs!\n" t) nil))))

(if (null pimacs-testing)
    (pimacs-greet-fn))

;; End of file.
