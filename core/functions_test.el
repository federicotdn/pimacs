;;; functions_test.el --- Tests for functions.go  -*- lexical-binding: t; -*-

(lt--deftest test-reverse ()
  (lt--should (equal (reverse '(1 2 3 4))
		     '(4 3 2 1))
	      "list reversed")
  (lt--should (equal (reverse ()) ()) "empty list reversed")
  (lt--should (equal (reverse [1 2 3 4])
		     [4 3 2 1])
	      "vec reversed")
  (lt--should (equal (reverse []) [])
	      "empty vec reversed"))

(lt--deftest test-eq ()
  (lt--should (eq 'a 'a) "a eq a")
  (lt--should-not (eq 'a 'b) "not a eq b")
  (lt--should-not (eq 1 1) "not 1 eq 1")
  (let ((foo "hello"))
    (lt--should (eq foo foo) "foo eq foo")))

(lt--deftest test-eql ()
  (lt--should (eql 'a 'a) "a eql a")
  (lt--should (eql 1 1) "1 eql 1")
  (lt--should-not (eql "hello" "hello") "not hello eql hello"))

(lt--deftest test-equal ()
  (lt--should (equal 'a 'a) "a equal a")
  (lt--should (equal 1 1) "1 equal 1")
  (lt--should (equal "hello" "hello") "hello equal hello"))

(lt--deftest test-delq ()
  (lt--should (equal '(1 2 3 4 5)
		     (delq 'a '(1 2 3 a 4 5)))
	      "list equal after delq (1)")
  (lt--should (equal '(1 2 3 4 5)
		     (delq 'a '(a 1 2 3 a 4 5)))
	      "list equal after delq (2)")
  (lt--should (equal '(1 2 3 4 5)
		     (delq 'a '(a 1 2 3 a 4 5 a a)))
	      "list equal after delq (3)")
  (lt--should (equal '()
		     (delq 'a '(a a a)))
	      "list equal after delq (4)")
  (lt--should (equal nil (delq 'a nil))
	      "list equal after delq (6)"))

(lt--deftest test-hash-table-eq ()
  (setq ht (make-hash-table :test 'eq))
  (setq key "key")
  (puthash 'a 123 ht)
  (puthash 100 "hello" ht)
  (puthash key "v" ht)

  (lt--should-not (gethash 'b ht))
  (lt--should-not (gethash 100 ht))
  (lt--should (eql (gethash 'a ht) 123) "eql 123")
  (lt--should (eql (gethash 'b ht 123) 123) "eql 123 default")
  (lt--should (equal (gethash key ht) "v"))
  (clrhash ht)
  (lt--should-not (gethash 'a ht)))
