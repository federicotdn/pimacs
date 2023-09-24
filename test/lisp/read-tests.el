;;; read-tests.el --- Tests for read.go  -*- lexical-binding: t; -*-

(lt--deftest test-unintern ()
  (lt--should (intern "hello") "intern a symbol ok")
  (lt--should (unintern "hello") "unintern a symbol ok")
  (setq mysym 'someverylongsymbol)
  (lt--should (unintern mysym) "unintern a symbol ok (2)")
  (lt--should-not (unintern "thisdoesnotexist") "not unintern a symbol ok"))

(lt--deftest test-read-integer ()
  (lt--should (equal #x10 16))
  (lt--should (equal #x0 0))
  (lt--should (equal #o10 8))
  (lt--should (equal #o0 0))
  (lt--should (equal #b110 6))
  (lt--should (equal #b0 0)))

(lt--deftest test-read-escape-hex ()
  (lt--should (equal ?\x10 16))
  (lt--should (equal ?\x0 0))
  (lt--should (equal ?\xA 10)))
