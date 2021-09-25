(progn
  (set 'welcome-message "Welcome to Pimacs!")

  (fset 'run-me (function (lambda () (prin1 "Hello!\n" t))))

  ;; (while t
  ;;   (set 'exp (read))
  ;;   (set 'res (eval exp))
  ;;   (princ res t)
  ;;   (princ "\n" t))

  )
