;;; character_table_test.el --- Tests for character_table.go  -*- lexical-binding: t; -*-

(lt--deftest test-char-table-parent ()
  ;; progn to prevent Emacs from using default indentation
  (progn
    (let ((child (make-char-table nil))
	  (parent (make-char-table nil)))
      (lt--should-not (char-table-parent child) "parent is nil")

      (set-char-table-parent child parent)
      (set-char-table-range parent '(140 . 141) "p")
      (lt--should (eq (char-table-parent child) parent) "parent is parent")

      (set-char-table-range child '(1 . 100) "foo")
      (lt--should (equal (char-table-range child 4) "foo") "val is foo")
      (lt--should (equal (char-table-range child 140) "p") "val is p"))))
