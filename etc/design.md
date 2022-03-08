# Design and general notes

This document contains many of the design choices and discoveries made while implementing Pimacs, in no particular order. It is useful for me in order to remember how Emacs itself works, and for other people that might be curious as to how Pimacs works.

Note that some of my assumptions of how Emacs works may be wrong or outdated. Feel free to suggest corrections in a pull request.

Here, I also try to compare the implementation differences between Emacs and Pimacs, and why they might be interesting.

## `execContext` and `lispObject`
The two central types to Pimacs' Elisp interpreter are `execContext` and `lispObject`.

The `execContext` structure contains all the state necessary to keep the interpreter working, such as the execution context stack (explained further on), the `obarray`, and all statically-defined symbols. The structure also contains other fields. `execContext` is defined in `exec_context.go`.

Similar to Emacs itself, the `lispObject` interface represents any valid Elisp object inside the interpreter. Only one function is required to fullfill the interface, with signature `getType() lispType` (i.e. the Elisp object must be able to provide a value representing its type). `lispType` is just an `int`. `lispObject` is defined in `types.go`.

## Subroutines
In Emacs, many Elisp functions are defined in C (such as `concat` or `+`). Special forms (such as `if` or `while`) are also defined in C. All of these together combined are referred to as "subroutines". The main difference between the two groups is that for special forms, arguments are not evaluated before being passed to them. This makes sense; for example, `if` needs to receive both the `then` and `else` expressions and decide which one to evaluate depending on the condition. If both expressions were evaulated before being handed to the `if`, that wouldn't be very useful! Apart from this, the mechanisms for implementing Elisp functions / special forms in C are quite similar.

So then, resumed: in Emacs Elisp functions and special forms are built using C functions plus a collection of C macros.

In Pimacs, Elisp functions and special forms are implemented using methods on the `execContext` structure, plus a manual call to a dedicated method for effectively registering the function or special form. The implementation method must return `(lispObject, error)`, and can accept either 0, 1, 2 all the way to 8 `lispObject` arguments, or otherwise `...lispObject`. These are then the valid signatures:
```go
type lispFn0 func() (lispObject, error)
type lispFn1 func(lispObject) (lispObject, error)
type lispFn2 func(lispObject, lispObject) (lispObject, error)
type lispFn3 func(lispObject, lispObject, lispObject) (lispObject, error)
type lispFn4 func(lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn5 func(lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn6 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn7 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFn8 func(lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject, lispObject) (lispObject, error)
type lispFnM func(...lispObject) (lispObject, error)
```
(8 was chosen as this is what Emacs uses as well)

Let's take a look at an example subroutine. Here's `car`:
```go
func (ec *execContext) car(obj lispObject) (lispObject, error) {
	if obj == ec.nil_ {
		return ec.nil_, nil
	}

	if !consp(obj) {
		return ec.wrongTypeArgument(ec.g.listp, obj)
	}
	return xCons(obj).car, nil
}
```

The subroutine must also be registered in the execution context, on startup:
```go
// Somewhere in the interpreter startup code
ec.defSubr1("car", ec.car, 1)
```

The arguments used are: `"car"` (name of the subroutine on the Elisp side), `ec.car` (the implementation method), and `1` (the minimum number of arguments this subroutine requires). `defSubr1` is used in particular because the method's signature accepts exactly 1 `lispObject`.

Note the usage of `ec.nil_` - this represents `nil` within the Elisp interpreter. It is not the Go `nil`!

Another example, here's `memq`:
```go
func (ec *execContext) memq(elt, list lispObject) (lispObject, error) {
	tail := list
	for ; consp(tail); tail = xCdr(tail) {
		if xCar(tail) == elt {
			return tail, nil
		}
	}

	if tail != ec.nil_ {
		return ec.wrongTypeArgument(ec.g.listp, list)
	}

	return ec.nil_, nil
}
```

And it is registered like so:
```go
ec.defSubr2("memq", ec.memq, 2)
```

## Stacks - Emacs
As the comment in the Emacs file `lisp.h` explains, Elisp uses three stacks for its runtime (in Emacs). It is useful to understand how they work, as Pimacs uses a slghtly different setup, but attempts to replicate the same functionality.

### C stack
The normal understanding of what "the stack" is in any C program. Function calls, stack-allocated variables, etc.

### `specpdl` stack
As the comment in `lisp.h` explains, "the `specpdl` stack keeps track of backtraces, unwind-protects and dynamic let-bindings". It is backed by an array.

Each entry of the `specpdl` stack is a `union` that may contain different fields depending on what the entry type is. The type easiest to understand is `SPECPDL_LET`. As the comment in the source says, it represents "a plain and simple dynamic let-binding".

For each entry type, there is specific code executed when the entry is pushed to the stack, and specific code executed when the entry is popped from the stack. In the case of pushing a `SPECPDL_LET` entry, first the specified symbol's current value is recorded in the entry itself, and then the symbol's value is set to the new one. When popping the entry, the symbol's value is restored to its original value. If a piece of code accesses the symbol's value between after the push but before the pop, it will receive the new value.

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

### Handler stack
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
      (throw 'bar 42))))
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
7. First, unwind the `specpdl` stack to the index specified in the entry. Then, use `longjmp` to restore the program environment back to when `catch 'bar` was called.
8. We are now back inside the call to `catch 'bar`. Remove the stack entry we added (`bar` tag) and all above it.
9. `catch` function returns normally (value: `42`).
10. We are now back inside the call to `catch 'foo`. Remove the stack entry we added (`foo` tag) and all above it. 
11. `catch` function returns normally (value: `nil`). Handler stack is now empty.

Note how the `baz` handler was removed by jumping back to the call to `catch 'bar` - it removed all elements up to and including the entry it itself added.

## Stacks - Pimacs
Pimacs only uses two stacks to implement the same behaviour as Emacs does.

### Go stack
The normal understanding of what the Go stack is (stack-allocated variables, function calls, etc.).

Go's ability to return multiple values in a function call is used to implement Elisp's non-local exits.

As mentioned before, in Pimacs, all subroutines must return `(lispObject, error)`. The usage of `error` here replaces Emacs' usage of `longjmp` for non-local exits. When a subroutine returns a non-nil `error`, other subroutines are expected to forward the `error` down the stack. This is, for better or for worse, a very common pattern in Go. A very specific set of subroutines - such as `catch` - can act on these `error` instances. This way, the implementation achieves something like what Emacs does, but following closely Go coding conventions.

So in one way, one could say that in Pimacs the Go stack replaces Emacs' handler stack. Let's take a look at the same code example as before, but now in Pimacs:
```elisp
(catch 'foo
  (catch 'bar
    (catch 'baz
      (throw 'bar 42))))
```

When we reach `throw`, the Go function call stack will look like this:
```
| <  call to throw   > |
| <call to catch 'baz> |
| <call to catch 'bar> |
| <call to catch 'foo> |
|______________________|
```

The implementation of `throw` is extremely simple:
```go
func (ec *execContext) throw(tag, value lispObject) (lispObject, error) {
    // ...
    return nil, &stackJumpTag{tag: tag, value: value}
}
```

And the implementation of `catch` boils down to:
```go
func (ec *execContext) catch(args lispObject) (lispObject, error) {
    tag, _ := ec.evalSub(xCar(args))
    obj, err := ec.progn(xCdr(args)) // Evaluate body

    if err != nil {
        // Somewhere up the stack, someone returned a non-nil error
        jmp, ok := err.(*stackJumpTag)
        
        if !ok {
            // The error does not come from a throw, just send it down the stack
            return nil, err
        }

        if jmp.tag == tag {
            // The error is a throw, and matches the tag for this catch!
            // Therefore, this catch evaluates to the value specified in the throw
            return jmp.value, nil
        }

        // Error is a throw, but tag does not match
        // Send it down the stack
        return nil, err
    }

    return obj, nil
}
```

So, from the Elisp example above, the call to `throw` will simply return a `(nil, &stackJumpTag{...})` pair, which will then be passed down the stack up to the call to `catch 'bar`. Then, that function call will simply return `(42, nil)` as a result - which is the expected behaviour!

### Execution context stack
As mentioned before, the `execContext` structure contains all the state belonging to the Elisp interpreter. One of the fields of this structure is called `stack` (the "execution context stack").

The execution context stack is an array of type `[]stackEntry`. Each entry can have a different type (that must implement `stackEntry`), and depending on the type each entry will contain different fields.

Entries with type `stackEntryLet` correspond, more or less, to Emacs' `SPECPDL_LET` stack entries. That is, they are used for dynamic let-bindings of symbol values.

Entries with type `stackEntryCatch` are used to record the presence of a call to `catch`. Interestingly though, these entries are not central to how the `catch`/`throw` mechanism works in Pimacs, because the mechanism itself is implemented using the methods' normal return value (see previous section). Why do these entries need to exist, then? The reason is because `throw` actually behaves differently depending on whether a `catch` has been set up or not. If there is no corresponding `catch`, `throw` will signal an error instead of doing the usual `longjmp` back to the corresponding `catch`. To do this, it needs to inspect the stack to see what the context is. In Emacs, this is done by looking in the handler stack. On Pimacs, this is done by looking in the execution context stack.

More entry types will be added as they become necessary.

In order to push entries to the execution context stack, methods such as `stackPushLet` and `stackPushCatch` should be used. However, the mechanism for popping stack entries is different when compared to Emacs' implementation for the `specpdl` stack. In Pimacs, Go's `defer` statement can be used in order to set up a hook that will automatically pop all execution context stack entries when the desired function returns. Here's an example that is more or less equivalent to the C one in the `specpdl` stack section above:

```go
func (ec *execContext) exampleFunction() (lispObject, error) {
    // foo's value is 42

    defer ec.unwind()()
    // Set up a function call that will unwind the exec context stack

    ec.stackPushLet(ec.g.foo, ec.t)
    // foo's value is now t

    return ec.innerFunction()
    // Before exampleFunction returns, the stack context entries
    // will be popped
}
```

Since the deferred function will execute for all return paths, this method makes it easy to push entries to the stack and then clean them up later, by only adding one line of code. The call to `unwind` returns an anonymous function that when called, will pop stack items back to the point where it was when the function itself was created (this is why `()()` is required).

If more fine-grained control of the stack is required, methods such as `stackSize` and `stackPopTo` can be used instead.

## Global variables
There are no global variables in the Pimacs code - everything is contained within the `execContext` structure. This should make it easier to understand the current state of the interpreter. It also allows for starting multiple interpreters under the same process or goroutine quite easily.

## Fatal error handling
When an error occurs - that the Pimacs developers did not foresee or handle correctly - a fatal error is raised. This is currently implemented using `panic`, which I think makes most sense in the contexts of the tools that Go offers.

Note that errors made by Pimacs _users_ are not considered to be fatal: for example, if the user executes `(+ "foo" 42)`, this will result in an error (technically, a call to `signal`) entirely within the interpreter. The top-level REPL function will handle this error (e.g. print it on screen), and allow the user to continue entering more code.

## Debugger
Pimacs does not have a debugger. This would probably be not very trivial to add, and I'm curious as to what changes would be necessary to make in order to add it. Maybe some of the assumptions I made when building the initial interpreter could be proven not to work when a debugger needs to be added.

## Macros and source code generation
Emacs C source code contains a lot of macro use. Some of the most interesting ones are defined in `lisp.h`. For example, `FOR_EACH_TAIL` allows for iterating through each tail of a list, and it even handles cyclical lists correctly by using a specific algorithm.

Some C source code is also generated by a separate program called `make-docfile`. During the build process, this program (whose source code is in `make-docfile.c`) is fed a set of Emacs C source files. The program then outputs the following:
- A file called `globals.h` with the actual declaration of all symbols, variables and subroutines that have been declared throughout the codebase. This file allows, for example, C code to use `Qfoo` to refer to the symbol `foo`. `Qfoo` will be a macro that returns a `Lisp_Object` representing the symbol. Note: this is not related to the `obarray`.
- A file called `etc/DOC`. See next section.
- Maybe more things, I am not sure.

Since Go does not have macros, and I want to avoid source code generation as much as possible (or completely), so Pimacs takes a slightly different approach. I attempted to make the process of declaring new subroutines, symbols and variables as least annoying as possible.

First, all statically defined symbols must be defined in the `lispGlobals` structure in `symbols.go`. "Declaring a subroutine" and "declaring a variable" just mean declaring a symbol and then assigning it a function value or data value, respectively (Elisp is a "Lisp-2"). The symbol must then be assigned a name via the `createSymbols` method function (the name visible from Elisp, that is). Lastly, in case of subroutines or variables, a call to `defSubrX` or `defVar` will be necessary. Errors can be declared as well in a similar fashion, using `putError` instead.

To resume:

In order to declare a symbol with no data or function value:
1. Add it to `lispGlobals`.
2. Initialize it in `createSymbols`.

In order to declare a variable:
1. Add it to `lispGlobals`.
2. Initialize it in `createSymbols`.
3. Call `defVar` somewhere during the interpreter initialization process.

In order to declare a subroutine (function or special form):
1. Create the corresponding Go method.
2. Call `defSubrX` somewhere during the interpreter initialization process.
3. If static access to the corresponding symbol is needed, add it to `lispGlobals` and initialize it in `createSymbols`. In most cases, static access is not needed.

In order to declare an error:
1. Add it to `lispGlobals`.
2. Initialize it in `createSymbols`.
3. Call `putError` somewhere in `errors.go`.

## Documentation strings
As mentioned in the previous section, somewhere in the Emacs build process, a file called `etc/DOC` is generated. This file contains the documentation of all variables and subroutines defined in C code (in C, they are written as comments - `make-docfile` picks those up). When Emacs starts up, it reads this file and then sets each element's documentation string to the indicated one (actually, it does this during the build process and not when starting up, but more on that later).

The process in Emacs then looks a bit like this:
1) `make-docfile` is run, which reads C source files.
2) `etc/DOC` is then generated will all documentation texts.
3) When `loadup.el` is loaded, it calls a function called `Snarf-documentation`.
4) `Snarf-documentation` (defined in `doc.c`) reads the documentation texts and sets them accordingly to each Lisp object.

In Pimacs, this functionality is not implemented yet. There are different ways this could be done:
- Use the same `etc/DOC` file Emacs generates and stick it into the Go binary using `go:embed`. During startup, read this file.
- Create a new property in the `subroutine` structure called `documentation`, and populate it somehow with a Go string. The problem with this is that there would be no easy way of keeping the method declarations and their corresponding documentation strings together (in the source file) in a neat way. Also, this would be kind of ugly because the documentation would be created as strings, not as Go comments.
- Parse the Go code somehow (a script or separate Go program?), picking up Go comments for each method, and generate something similar to `etc/DOC`. Use `go:embed` to stick this file into the Go binary and read it on interpreter startup. The (big) advantage here is that normal Go comments would be used to document methods. The disadvantage would be making the build process slightly more complicated. If I had to implement this option, I would ensure Pimacs could be built without the `etc/DOC` file actually existing or it being empty - it would just skip loading documentation strings, in those cases.

## Build process
Required reading: [Building Emacs](https://www.gnu.org/software/emacs/manual/html_node/elisp/Building-Emacs.html).

TL;DR: When compiling Emacs, the first step is to actually generate a bare Emacs binary called `temacs`. This is then run, all Lisp files and other resources are loaded, and then the program's state is dumped - somehow - into a file. The final `emacs` binary is a combination of `temacs` plus this dumped file. This is done to avoid Emacs having to load all Lisp files and resources from scratch during each startup.

The build process for Pimacs is currently just `go build`. This produces a single static binary. Ideally, this should remain like this - if I had to include additional resources or files, I would use `go:embed` to store them directly in the binary file. I have already experimented with this by storing a single `base.el` file inside the Pimacs binary, and then reading it during startup as if it were a normal Elisp source file. Also, the concept of the build process just being "Go code -> compiler -> binary" would be nice to maintain.

## Bytecode
Pimacs does not have a bytecode interpreter, but it would be interesting to add one. Unsure if the bytecode format should be the same as Emacs' (it would definitely make many things simpler).

## Native Elisp compilation
Some context [here](https://old.reddit.com/r/emacs/comments/myej3z/the_nativecompilation_branch_was_just_merged_into/). Not sure what the analogous would be in Go.

## Garbage collection
Being written in Go, Pimacs uses Go's garbage collector. This massively simplifies a lot of the implementation work for the interpreter. The Go garbage collector is also extremely battle-tested and is obviously much better than anything I could've coded myself.

## Concurrency with goroutines & channels
Given that all the interpreter state is inside `execContext`, it is quite easy to just start a new goroutine using `go`, create a new `execContext` there, and then start evaluating Elisp code on that separate goroutine.

I have experimented with creating an Elisp function called `pimacs-go` that takes a `BODY` argument, and evaluates it on a new goroutine. Together with that, I added a new Elisp data type that is backed by a Go channel, together with two functions: `pimacs-go-channel-send` and `pimacs-go-channel-receive` (corresponding to `ch<-` and `<-ch` respectively). The implementation can be found around [this file](https://github.com/federicotdn/pimacs/blob/4488276ecbad2de75f67b1b798380dc9d48eddb2/elisp/subroutines.go), but I've since removed it since it was too janky and improvised. However, this would be interesting to re-visit in the future, since it is one of Go's strongest aspects, and it would be very useful to make use of it from the Elisp side.

Some difficult questions for this topic are still to be answered, such as how much state should `execContext` instances in different goroutines share.

## Helper functions
The `helpers.go` file contains a collection of functions that make writing Pimacs code a bit easier.

In Emacs, these would be the `CONSP`, `XCONS`, `XCAR`, (...) macros.

In Pimacs, these are simply Go functions that can be called for checking the type of a `lispObject`, or for accessing attributes of specific types (e.g. the function value of a symbol). It is worth noting that these helpers are not methods on `execContext`, but rather simple functions, so they can be called without needing to use `ec.<name>`.

Some of the helper functions' names begin with an `x`. This means that when calling the helper, **all arguments must meet some preconditions or else the helper may `panic`.** This is important because it means the Pimacs programmer must carry out these checks in order to use the helpers correctly. A `panic` is not a situation Pimacs can recover from (for now, at least). Therefore, a `panic` resulting from a helper call implies a wrong assumption that the Pimacs programmer has made.

Some examples. We can look at the predicate `symbolp` helper, and `xSymbol` as well. The following code is not correct, as it may `panic` depending on what type of Lisp object `obj` actually is:
```go
obj, _ := ec.someFunction()

// Set the symbol's data value to t
xSymbol(obj).value = ec.t

// But if obj is not a symbol, the above will panic!
```

This would be the correct, equivalent code:
```go
obj, _ := ec.someFunction()

if symbolp(obj) {
    // Set the symbol's data value to t
    xSymbol(obj).value = ec.t
}
```

Other helpers include:
- `xEnsure`: Takes in the result of calling a subroutine (i.e. `(lispObject, error)`) and returns only the `lispObject` - discarding the error only if it is `nil`. Otherwise, it `panic`s.
- `xErrOnly`: The opposite of `xEnsure`. Only returns the `error`, if the `lispObject` is `nil` (Go `nil`, not Elisp `nil`). If the `lispObject` is not `nil`, it `panic`s.

## Project structure
Currently, the project structure is as similar to Emacs' as possible, in order to make porting of code easier. Some names have been slightly modified though; for example `alloc.c` is `allocation.go`, and `fns.c` is `functions.go`.

## User interface
If the project ever reaches enough maturity in order for it to have a user interface, it would be interesting to take an approach similar to [xi-editor's](https://github.com/xi-editor/xi-editor). This would mean that all Elisp functions relating to manipulation of the UI would actually be communicating via some protocol with a frontend. The frontend could be implemented in different ways, e.g. TUI, native (GTK, Qt) or Web technologies.

## Compatibility with Emacs
The current main design guide for Pimacs is to implement an Elisp interpreter that allows for running any Elisp script without modifications. This means that Pimacs needs to implement functions, special forms and variables using Emacs' current implementation as a guide. The version of Emacs in question is, for the moment, just whatever is on the `master` branch. This is not ideal since this implies the target implementation is constantly moving.

If Pimacs adds any extra functionality, it is namespaced with `pimacs-` or `pimacs--` in order to make the distinction with the Emacs codebase clearer.

## Testing
A very basic set of tests are implemented in `interpreter_test.go`. However, these tests are a bit of a hassle to write, since all Elisp code must be written in strings, as well as the desired outputs.

Ideally, Pimacs should be able to load Emacs' [ERT](https://www.gnu.org/software/emacs/manual/html_node/ert/), and then just use the same set of Elisp test files.
