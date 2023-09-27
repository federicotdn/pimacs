;;; data-tests.el --- Tests for data.go  -*- lexical-binding: t; -*-

(lt--deftest test-1+ ()
  (lt--should (equal (1+ 1) 2)))
