;;; data-tests.el --- Tests for data.go  -*- lexical-binding: t; -*-

(lt--deftest test-1+ ()
  (lt--should (equal (1+ 1) 2)))

(lt--deftest test-arithmetic-op ()
  (lt--should (= (+ 1 2 3) 6))
  (lt--should (= (+ 1) 1))
  (lt--should (= (+) 0))

  )

(lt--deftest test-arithmetic-cmp ()
  (lt--should (= 1 1))
  (lt--should-not (= 1 2))

  (lt--should (/= 1 2))
  (lt--should-not (/= 1 1))

  (lt--should (< 1 2))
  (lt--should-not (< 2 1))

  (lt--should (> 2 1))
  (lt--should-not (> 1 2))

  (lt--should (<= 1 1))
  (lt--should (<= 1 2))
  (lt--should-not (<= 3 1))

  (lt--should (>= 1 1))
  (lt--should (>= 2 1))
  (lt--should-not (>= 1 3))

  (lt--should (< 1 2 3 4))
  (lt--should-not (< 1 2 3 4 3))

  (lt--should (< 1))
  (lt--should (= 100)))
