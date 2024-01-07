;;; read-tests.el --- Tests for read.go  -*- lexical-binding: t; -*-

(lt--deftest test-unintern ()
  (lt--should (intern "hello") "intern a symbol ok")
  (lt--should (unintern "hello") "unintern a symbol ok")
  (setq mysym 'someverylongsymbol)
  (lt--should (unintern mysym) "unintern a symbol ok (2)")
  (lt--should-not (unintern "thisdoesnotexist") "not unintern a symbol ok"))

(lt--deftest test-read-integer ()
  (lt--should (equal #x10 16))
  (lt--should (equal #x010 16))
  (lt--should (equal #x0 0))
  (lt--should (equal #o10 8))
  (lt--should (equal #o0 0))
  (lt--should (equal #b110 6))
  (lt--should (equal #b0 0)))

(lt--deftest test-read-char-escape-hex ()
  (lt--should (equal ?\x10 16))
  (lt--should (equal ?\x0 0))
  (lt--should (equal ?\xA 10))
  (lt--should (equal ?\x0A 10))
  (lt--should (equal ?\x00A 10)))

(lt--deftest test-read-char-escape-mod ()
  (lt--should (equal ?\M-a #o1000000141) "M-a")
  (lt--should (equal ?\M-v #o1000000166) "M-v")
  (lt--should (equal ?\C-g #o7) "C-g")
  (lt--should (equal ?\M-\C-g #o1000000007) "M-C-g")
  (lt--should (equal ?\C-? #o177) "C-?")
  ;; Reading ?\C- alone yields -1, but using it in a larger expression
  ;; yields 67108896 (Emacs bug?).
  (lt--should (equal (car (read-from-string "?\\C-")) -1) "C-")
  ;; Note: at time of writing, Emacs also requires a space after the
  ;; cdr here, otherwise it results in a parsing error (probably a bug).
  (lt--should (equal (cons ?\C- ?\C- ) (cons 67108896 67108896))))
