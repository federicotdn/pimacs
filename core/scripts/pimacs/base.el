(progn
  (fset 'pimacs-greet-fn (function (lambda () (progn (princ "Welcome to Pimacs!\n" t) nil))))
  (if (null pimacs-testing)
      (pimacs-greet-fn)))
