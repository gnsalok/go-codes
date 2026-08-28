# Lesson 21: Closures

## Learning goal

By the end of this lesson, you should be able to explain what a closure is,
create one in Go, and use it to keep private state between function calls.

## What is a closure?

A closure is a function value that refers to variables declared outside the
function's own body. The function keeps access to those variables for as long
as the function can still be called.

In this lesson, `NewCounter` creates a local variable named `count` and returns
an anonymous function:

```go
func NewCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
```

The returned function is a closure because it captures `count`. Although
`NewCounter` finishes, `count` remains available to the returned function.

## Counter example

```go
counter := NewCounter()

fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3
```

Here is what happens step by step:

1. `NewCounter()` creates `count` with a value of `0`.
2. It returns a function that captures that particular `count` variable.
3. The first call increments `count` to `1` and returns it.
4. Later calls use the same captured variable, producing `2`, then `3`.

The state is private: code that receives the closure can change `count` only by
calling the returned function.

## Independent closures

Every call to `NewCounter` creates a new captured variable:

```go
first := NewCounter()
second := NewCounter()

fmt.Println(first())  // 1
fmt.Println(first())  // 2
fmt.Println(second()) // 1
```

`first` and `second` do not share state. Each closure has its own `count`.

## Real-world example: HTTP timing middleware

Closures are commonly used to build HTTP middleware. Middleware wraps an
existing handler and adds behavior such as logging, authentication, metrics, or
panic recovery.

The following middleware measures how long an HTTP handler takes to run:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Run the handler wrapped by this middleware.
		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s completed in %s",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	hello := http.HandlerFunc(helloHandler)
	http.Handle("/hello", timingMiddleware(hello))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### When does each function run?

This line is evaluated during application startup:

```go
http.Handle("/hello", timingMiddleware(hello))
```

At that point, `timingMiddleware(hello)` captures `hello` in the parameter
named `next`. It returns a new handler, and `http.Handle` registers that new
handler for `/hello`. The original `helloHandler` does not run yet.

When a request arrives, the HTTP server calls the returned handler. The
execution order is:

```text
GET /hello
    -> returned middleware handler starts the timer
    -> next.ServeHTTP(w, r)
    -> hello.ServeHTTP(w, r)
    -> helloHandler(w, r)
    -> middleware logs the elapsed time
```

The original handler is called at this exact line:

```go
next.ServeHTTP(w, r)
```

`next` has the `http.Handler` interface type. In this example, its concrete
value is `hello`, which is an `http.HandlerFunc`. The `HandlerFunc` adapter's
`ServeHTTP` method calls the underlying function. Conceptually, it behaves like
this:

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

The closure is useful here because the returned handler remembers `next` after
`timingMiddleware` has returned. The same pattern can capture other middleware
configuration, such as a logger, an authentication service, or a rate limit.

## Run the lesson

From the repository root:

```bash
go run ./lessons/21_closures
```

Expected output:

```text
First counter: 1
First counter: 2
First counter: 3
Second counter: 1
```

## Run the tests

```bash
go test ./lessons/21_closures
```

The tests verify that a closure remembers its state and that separately created
closures remain independent.

## Exercise

Change `NewCounter` so it accepts a starting value:

```go
counter := NewCounter(10)

fmt.Println(counter()) // 11
fmt.Println(counter()) // 12
```

For an extra challenge, create a closure that accepts an increment amount, so
the caller can increase or decrease the counter on each call.

## Key points

- Functions are values in Go, so they can be returned from other functions.
- A closure captures variables from its surrounding scope.
- Captured variables keep their values between calls.
- Separate closures can maintain separate private state.
- HTTP middleware uses closures to remember the next handler and configuration.
