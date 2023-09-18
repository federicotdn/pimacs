;;; read_test.el --- Tests for read.go  -*- lexical-binding: t; -*-

(lt--deftest test-unintern ()
  (lt--should (intern "hello") "intern a symbol ok")
  (lt--should (unintern "hello") "unintern a symbol ok")
  (setq mysym 'someverylongsymbol)
  (lt--should (unintern mysym) "unintern a symbol ok (2)")
  (lt--should-not (unintern "thisdoesnotexist") "not unintern a symbol ok"))
