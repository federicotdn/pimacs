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

Try out a nonlocal exit:
```elisp
> (catch 'test (throw 'test 123))
123
```

## Design and general notes
This section contains many of the design choices and discoveries made while implementing Pimacs, in no particular order. Note that some of my assumptions of how Emacs works may be wrong or outdated. Feel free to suggest corrections in a pull request.

### Stacks - Emacs
As the comment in the Emacs file `lisp.h` explains, Emacs Lisp uses three stacks for its runtime.

#### C stack
The normal understanding of what "the stack" is in any C program. Function calls, stack-allocated variables, etc.

#### `specpdl` stack

#### Handler stack
This stack is implemented using a doubly-linked list, and it keeps track of the handlers needed to implement the `catch`/`throw` and `condition-case`/`signal` mechanisms (non-local exits).

Each element of the handler stack is a structure with a set of fields. One of those fields, of type `sys_jmp_buf` and called `jmp`, is used to actually record the calling environment so that it can be 'restored' later. Here's a simplified version of the data structure used:
```c
enum handlertype { CATCHER, CONDITION_CASE, CATCHER_ALL };

struct handler
{
  // Handler type
  enum handlertype type;

  // Tag
  Lisp_Object tag_or_ch;

  // Linked list links
  struct handler *next;
  struct handler *nextfree;

  // Environment to restore when popping here
  sys_jmp_buf jmp;

  // Index of specpdl when this entry was added
  specpdl_ref pdlcount;

  // (Some fields omitted)
};
```

And then, a pointer is used to point to the top of the stack:
```c
struct handler *m_handlerlist;
```

How this is used: for example, the `catch` function pushes a new entry to the handler stack, by calling `setjmp` and storing the associated environment in the stack entry's `jmp`. The tag used to call `catch` is also stored in the stack entry.

Later, when `throw` is called, the stack is iterated top-to-bottom, looking for an entry with a matching tag. When one is found, its `jmp` property is used to call `longjmp`, thus restoring the program environment to when the matching `catch` was called. The handler stack itself is also restored to the configuration it had before `catch` was called - whether a corresponding `throw` was called or not.

In diagram form, let's say we have the following code:
```elisp
(catch 'foo
  (catch 'bar
    (catch 'baz
      (throw 'foo nil))))
```

When the `throw` is about to be called, the handler stack would then look like this:

```
| { handler for 'baz } |
| { handler for 'bar } |
| { handler for 'foo } |
|______________________|
```

Something important to note is that when an entry is added to the handler stack, the entry will also contain the current `specpdl` index (`pdlcount` field). When we pop back through the handler stack either with `throw` or `signal`, the `specpdl` stack entries are also popped as well, until the index specified in the handler stack entry. In other words, the `specpdl` stack is also restored to the state it was in when the corresponding `catch` or `condition-case` was called.

See: [`setjmp` and `longjmp`](https://en.wikipedia.org/wiki/Setjmp.h) for more information.

### Stacks - Pimacs

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

### Compatibility with Emacs
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

## License
Like Emacs, Pimacs is licensed under the GNU General Public License, version 3.
