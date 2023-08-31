<p align="center">
  <img alt="pimacs" src="https://github.com/federicotdn/pimacs/raw/main/extra/logo.png" width="50%">
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
You can also use `./pimacs --load myfile.el` to execute an Elisp script.

Note that many, many Elisp functions and macros are **not** implemented. You can, however, use the following (among others):
```
if while eval funcall apply quote function progn
unwind-protect condition-case throw signal
prin1 prin1-to-string print princ
intern read-from-string read
length equal assq memq get put plist-get plist-put nconc
null symbol-plist symbol-name set fset eq defalias
setcar setcdr car cdr cons list + <
current-buffer get-buffer set-buffer buffer-string insert
```
(and additionally, many predicate functions e.g. `stringp`, `consp`, etc.)

Note that some of these may be only partially implemented.

### Examples
Set a variable and read it:
```elisp
> (set 'greeting "hello")
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
> (set 'vec [1 2 3 4])
[1 2 3 4]
> (vectorp vec)
t
```

## Design and general notes
In order to read about the design choices for Pimacs, how it works internally, and how it is different from Emacs' Elisp interpreter, see the [design.md](extra/design.md) document.

## Tests
Use `make test` to run the test suite.

## License
Like Emacs, Pimacs is licensed under the GNU General Public License, version 3.
