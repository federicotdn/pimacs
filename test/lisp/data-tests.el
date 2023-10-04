;;; data-tests.el --- Tests for data.go  -*- lexical-binding: t; -*-

(lt--deftest test-1+ ()
  (lt--should (equal (1+ 1) 2)))

(lt--deftest test-arithmetic-op ()
  (lt--should (= (+ 1 2 3) 6))
  (lt--should (= (+ 1) 1))
  (lt--should (= (+) 0))

  (lt--should (= (logior 4 2 1) 7)))

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

(lt--deftest test-aref ()
  (let ((v [100 200 300])
	(s "hello")
	(s2 "aaább"))
    (lt--should (= (aref v 0) 100))
    (lt--should (= (aref s 0) ?h))
    (lt--should (= (aref s 2) ?l))
    (lt--should (= (aref s2 2) ?á))
    (lt--should (= (aref s2 4) ?b))))

(lt--deftest test-aset ()
  (let ((v [100 200 300]))
    (aset v 0 99)
    (lt--should (= (aref v 0) 99))))
