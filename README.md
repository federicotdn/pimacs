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

Try out a non-local exit:
```elisp
> (catch 'test (throw 'test 123))
123
```

## Design and general notes
This section contains many of the design choices and discoveries made while implementing Pimacs, in no particular order. Note that some of my assumptions of how Emacs works may be wrong or outdated. Feel free to suggest corrections in a pull request.

Here, I also try to compare the implementation differences between Emacs and Pimacs, and why they might be interesting.

### Stacks - Emacs
As the comment in the Emacs file `lisp.h` explains, Emacs Lisp uses three stacks for its runtime (in Emacs). It is useful to understand how they work, as Pimacs uses a slghtly different setup, but attempts to replicate the same functionality.

#### C stack
The normal understanding of what "the stack" is in any C program. Function calls, stack-allocated variables, etc.

#### `specpdl` stack
As the comment in `lisp.h` explains, "the `specpdl` stack keeps track of backtraces, unwind-protects and dynamic let-bindings". It is backed by an array.

Each entry of the `specpdl` stack is a `union` that may contain different fields depending on what the entry type is. The type easiest to understand is `SPECPDL_LET`. As the comment in the source says, it represents "a plain and simple dynamic let-binding".

For each entry type, there are is specific code executed when the entry is pushed to the stack, and specific code executed when the entry is popped from the stack. In the case of pushing a `SPECPDL_LET` entry, first the specified symbol's current value is recorded in the entry itself, and then the symbol's value is set to the new one. When popping the entry, the symbol's value is restored to its original value. If a piece of code accesses the symbol's value between after the push but before the pop, it will receive the new value.

Below is an example of the `specpdl` stack being used to bind the (dynamic) value of a symbol `foo`.
```c
void example_function ()
{
  specpdl_ref count = SPECPDL_INDEX ();

  // Symbol foo's value is 42 (for example)

  specbind (Qfoo, Qt);
  // Now, foo's value is t

  // For any code executed by inner_function,
  // foo's value will be t
  inner_function ();

  unbind_to (count, Qnil);
  // foo's value is now 42 again (i.e. the
  // original value)
}
```

At the beginning of `example_function`, the current `specpdl` stack index is stored so that we can unwind it back to the same point again when the function ends.
Then, `specbind` is called to set the value of symbol `foo` to `t`. Internally, `specbind` will push a `SPECPDL_LET` entry to the stack. This entry contains the symbol's original value. Then, it will set the symbol's value to the new one (`t`, in this case).
Now, any code that accesses `foo`'s value will receive `t` as a result. This includes, in this example, `inner_function`, or any functions called by it.
At the end of `example_function`, we undo the let-binding by calling `unbind_to`. For this to work, we need to specify up to what point of the `specpdl` stack entries should be popped. Internally, `unbind_to` will execute the cleanup code associated with each entry type. As mentioned before, for `SPECPDL_LET` this involves restoring the symbol's original value.

It is important to note that it is the C programmer's responsibility to ensure that the `specpdl` stack is restored to its original state at the end of the function or block. As far as I'm aware there is no mechanism that assists the programmer in ensuring they remember this.

#### Handler stack
This stack is implemented using a doubly-linked list, and it keeps track of the information needed to implement the `catch`/`throw` and `condition-case`/`signal` mechanisms (non-local exits).

Each entry of the handler stack is a data structure with a set of fields. Here's a simplified version of the structure used:
```c
enum handlertype { CATCHER, CONDITION_CASE /* ... (Some values omitted) */ };

struct handler
{
  // Handler type:
  // Most importantly, distinguishes between entries set up
  // by `catch` vs. entries set up by `condition-case`.
  enum handlertype type;

  // Tag:
  // When using `catch`, specifies the tag used.
  Lisp_Object tag_or_ch;

  // Environment:
  // This is where Emacs actually stores the program environment
  // so that it can be restored later on. It is populated using
  // the `setjmp` function, and restored using `longjmp`.
  sys_jmp_buf jmp;

  // specpdl index:
  // When a new handler stack entry is added, the current specpdl
  // index is recorded as well.
  specpdl_ref pdlcount;

  // ... (Some fields omitted)
};
```

See [`setjmp` and `longjmp`](https://en.wikipedia.org/wiki/Setjmp.h) for more information on how these two functions work.

Finally, a pointer is used to point to the top of the stack. This is the way Emacs accesses the handler stack during runtime.
```c
struct handler *m_handlerlist;
```

How is the handler stack actually used? An example: when called, the `catch` function pushes a new entry to the handler stack. The fields of the new entry are populated like so:
- `type` = `CATCHER`
- `tag_or_ch` = value of the tag (e.g. `'done`)
- `jmp` = value set by calling `setjmp`
- `pdlcount` = value of the current `specpdl` stack index
- ... (some fields ommitted)

Later, when `throw` is finally called, the stack is iterated top-to-bottom, looking for an entry with a matching tag (`tag_or_ch`). When one is found, its `jmp` property is used to call `longjmp`, thus restoring the program environment to when the matching `catch` was called. At the same time, the `pdlcount` property is used to unwind the `specpdl` stack back to the same configuration it had when `catch` was called. The handler stack itself is also restored to the configuration it had before `catch` was called - all of this happens whether a corresponding `throw` was called or not.

A more specific example, in diagram form. Let's say we have the following code:
```elisp
(catch 'foo
  (catch 'bar
    (catch 'baz
      (throw 'bar nil))))
```

When `throw` is about to be called, the handler stack would then look like this:

```
| { handler for 'baz } |
| { handler for 'bar } |
| { handler for 'foo } |
|______________________|
```

The sequence of events before and after `throw` is called would be (starting with an empty handler stack):
1. `catch 'foo` is called. An entry is pushed to the handler stack (`foo` tag).
2. `catch 'bar` is called. An entry is pushed to the handler stack (`bar` tag).
3. `catch 'baz` is called. An entry is pushed to the handler stack (`baz` tag).
4. `throw 'bar` is called.
5. We iterate through the stack top-to-bottom until we find a handler with `bar`.
6. The handler entry is found (#2).
7. Use `longjmp` to restore the program environment back to when `catch 'bar` was called.
8. We are now back inside the call to `catch 'bar`. Remove the stack entry we added (`bar` tag) and all above it.
9. `catch` function returns normally.
10. We are now back inside the call to `catch 'foo`. Remove the stack entry we added (`foo` tag) and all above it. 
11. `catch` function returns normally. Handler stack is now empty.

Note how the `baz` handler was removed by jumping back to the call to `catch 'bar` - it removed all elements up to and including the entry it itself added.

### Stacks - Pimacs
#### Go stack
Todo.

#### `execContext` stack
Todo. (`defer` of `stackPopTo`)

### The `lispObject` type
Todo.

### Fatal error handling
Todo. (panics)

### Registering subroutines

### Macros and source code generation
Todo. (`make-docfile`, `globals.h`)

### Documentation strings
Todo. (`go:embed` + gen script, const strings)

### Garbage collection
Todo.

### Concurrency with goroutines
Todo. (multiple interpreter instances, channels)

### Usage of `go:embed`
Todo.

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
