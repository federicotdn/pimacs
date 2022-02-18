<p align="center">
  <img alt="pimacs" src="https://github.com/federicotdn/pimacs/raw/main/etc/logo.png" width="50%">
  <br/>
</p>

## Summary

A very incomplete and experimental implementation of an Emacs Lisp interpreter, written in Go.

## Project goals
- Practice Go development.
- Learn more about the Emacs internals, particularly the Emacs Lisp interpreter.
- Learn about Lisp interpreter design in general.
- Practice reading C code.
- Learn about lower level functions such as `setjmp`, `longjmp`, etc.

## Usage
Assuming you have the Go compiler installed, simply use `make build` to compile Pimacs, and then `./pimacs` to start the REPL.
You can also use `./pimacs --load myfile.el` to execute an Emacs Lisp script.

Note that many, many Emacs Lisp functions and macros are **not** implemented. You can, however, use the following (among others):
```
if while eval funcall apply quote function progn
unwind-protect condition-case throw signal
prin1 prin1-to-string print princ
intern read-from-string read
length equal assq memq get put plist-get plist-put
null car cdr symbol-plist symbol-name set fset eq defalias
cons list + <
```

Note that some of these may be only partially implemented.

### Examples
Set a variable and read it:
```
> (set 'greeting "hello")
"hello"
> greeting
"hello"
```

Create a function and call it:
```
> (fset 'twice (function (lambda (x) (+ x x))))
(lambda (x) (+ x x))
> (twice 21)
42
```

Try out a nonlocal exit:
```
> (catch 'test (throw 'test 123))
123
```

## Design and general notes
This section contains many of the design choices and discoveries made while implementing Pimacs, in no particular order. Note that some of my assumptions of how GNU Emacs works may be wrong or outdated.

### Stacks
Todo. (See `lisp.h` comment on C stack, `specpdl` and handler stack, Go `error` return value)

### Macros and source code generation
Todo. (`make-docfile`, `globals.h`)

### Garbage collection
Todo.

### Concurrency with goroutines
Todo. (multiple interpreter instances, channels)

### Usage of `go:embed`
Todo.

### Fatal error handling
Todo. (panics)

### Helper functions
Todo. (`xCdr` and friends)

### User interface
Todo.

### Compatibility with GNU Emacs
Todo. (pimacs extensions `pimacs-*`)

### Build system
Todo. (single binary result, build tools)

## Tests
Use `make test` to run the test suite.

## Comment Tags
Since there is so much code to write, some functions contain tags within comments that specify what is missing or what should be done in order to consider the function implementation as complete.

- `incomplete`: Code implements some functionality correctly, but not 100% of what it should.
- `revise`: Code implements the entire functionality, but needs to be revised or re-thought.
- `errors`: Code is missing proper error handling.
- `stub`: Code does not implement any functionality yet.
- `tests`: Code needs to be tested using unit tests, or more tests need to be added for it.
