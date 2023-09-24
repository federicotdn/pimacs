;;; allocation-tests.el --- Tests for allocation.go  -*- lexical-binding: t; -*-

(lt--deftest test-make-vector ()
  (lt--should (equal (make-vector 5 1) [1 1 1 1 1]))
  (lt--should (equal (make-vector 0 0) [])))
