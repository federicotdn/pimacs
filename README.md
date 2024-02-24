<p align="center">
  <img alt="pimacs" src="https://github.com/federicotdn/pimacs/raw/main/etc/logo.png" width="50%">
  <br/>
</p>

## Summary

A partial, experimental implementation of an Elisp (Emacs Lisp) interpreter, written in Go.

## Project goals
- Practice Go development.
- Learn more about the Emacs internals, particularly the Elisp interpreter.
- Learn about Lisp interpreter design in general.
- Practice reading C code.
- Learn about lower level functions such as `setjmp`, `longjmp`, etc.
- Learn about the general challenges found when creating a text editor.

## Usage
Assuming you have the Go compiler installed, simply use `make build` to compile Pimacs, and then `./pimacs` to start the REPL.

Note that many, many Elisp functions and macros are **not** implemented. You can, however, use the following (among others):
```
intern read-from-string read load eval funcall apply progn prog1 cond
setq and or if while quote function defconst let let* catch
unwind-protect condition-case throw signal prin1 print princ
prin1-to-string null sequencep consp listp symbolp stringp
number-or-marker-p char-or-string-p integerp numberp bufferp
characterp char-table-p vectorp boundp fboundp makunbound fmakunbound
car cdr car-safe cdr-safe setcar setcdr symbol-plist symbol-name set
fset symbol-value symbol-function eq defalias + < bare-symbol cons
list length equal eql assq assoc memq get put plistp plist-get plist-put
nconc provide nreverse reverse require nthcdr nth mapcar buffer-string
insert current-buffer set-buffer get-buffer buffer-name buffer-list
get-buffer-create read-from-minibuffer getenv-internal recursive-edit
make-char-table char-table-range set-char-table-range
char-table-parent set-char-table-parent multibyte-string-p %
```

Note that some of these may be only partially implemented, or be a stub/placeholder.

### Examples
Set a variable and read it:
```elisp
> (setq greeting "hello")
"hello"
> greeting
"hello"
```

Create a function and call it:
```elisp
> (fset 'twice (function (lambda (x) (+ x x))))
(lambda (x) (+ x x))
> (twice 21)
42
```

Try out a non-local exit:
```elisp
> (catch 'test (throw 'test 123))
123
```

Create a vector of integers:
```elisp
> (setq vec [1 2 3 4])
[1 2 3 4]
> (vectorp vec)
t
```

Use backquotes:
```elisp
> (setq lst '(2 3 4))
(2 3 4)
> `(1 ,@lst 5 6)
(1 2 3 4 5 6)
```

Create a multibyte string:
```elisp
> (setq s "ñandú")
"ñandú"
> (multibyte-string-p s)
t
```

## Design and general notes
In order to read about the design choices for Pimacs, how it works internally, and how it is different from Emacs' Elisp interpreter, see the [design.md](etc/design.md) document.

## Tests
Use `make test` to run the test suite.

## Similar projects
Check out [Rune](https://github.com/CeleritasCelery/rune), a re-implementation of Emacs from scratch using Rust.

## License
Like Emacs, Pimacs is licensed under the GNU General Public License, version 3.
