(progn
  (set 'pimacs-welcome-message "Welcome to Pimacs!")
  (fset 'pimacs-greet-fn (function (lambda () (progn (princ "Hello!\n" t) nil)))))
