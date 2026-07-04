# Errors in Go

## 1. The Basics: The `error` Interface

In Go, `error` is a built-in interface. Any type that implements a single `Error()` method returning a string can be used as an error.

```go
type error interface {
    Error() string
}

```

### Basic Usage Example

When a function can fail, you return a tuple: `(result, error)`. If everything goes well, you return `nil` for the error.

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Creating a simple error using the errors package
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err) // Output: Error occurred: cannot divide by zero
		return
	}
	fmt.Println("Result:", result)
}

```

---

## 2. Formatting Errors with `fmt.Errorf`

If you need to add dynamic data to your error message (like an ID or a variable state), use `fmt.Errorf`.

```go
func checkUsername(username string) error {
    if len(username) < 3 {
        return fmt.Errorf("username %q is too short (min 3 chars)", username)
    }
    return nil
}

```

---

## 3. Custom Error Types

Sometimes a simple string isn't enough. You might want to attach extra context to the error (like an HTTP status code or a machine-readable error code). You can achieve this by creating a custom struct that implements the `error` interface.

```go
package main

import "fmt"

// 1. Define the custom struct
type QueryError struct {
	Query string
	Err   error
}

// 2. Implement the error interface
func (e *QueryError) Error() string {
	return fmt.Sprintf("query %q failed: %v", e.Query, e.Err)
}

func executeQuery() error {
	return &QueryError{
		Query: "SELECT * FROM users",
		Err:   fmt.Errorf("database connection lost"),
	}
}

func main() {
	err := executeQuery()
	if err != nil {
		fmt.Println(err) // Output: query "SELECT * FROM users" failed: database connection lost
	}
}

```

---

## 4. Modern Error Handling: Wrapping, Unwrapping, `errors.Is`, and `errors.As` (Go 1.13+)

When an error passes up through multiple layers of your application, you often want to add context to it without losing the original error. Go allows you to **wrap** errors using `%w` in `fmt.Errorf`.

### Wrapping Errors with `fmt.Errorf` and `%w`

Wrapping means: add extra context to an error, but keep the original error available for later checks.

```go
package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("file does not exist")
}

func startApp() error {
	err := readConfig()
	if err != nil {
		// %w wraps the original error.
		return fmt.Errorf("could not start app: %w", err)
	}
	return nil
}

func main() {
	err := startApp()
	fmt.Println(err)
	// Output: could not start app: file does not exist
}

```

Short form:

```go
originalErr := errors.New("permission denied")

// Wrapping the original error with more context
wrappedErr := fmt.Errorf("failed to delete file: %w", originalErr)

```

Use `%w` only when you want the caller to be able to inspect the original error with `errors.Is`, `errors.As`, or `errors.Unwrap`. If you only want text formatting and do not want wrapping, use `%v`.

### Unwrapping Errors with `errors.Unwrap`

`errors.Unwrap` returns the next error inside a wrapped error chain.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	rootErr := errors.New("connection refused")
	wrappedErr := fmt.Errorf("database failed: %w", rootErr)

	unwrappedErr := errors.Unwrap(wrappedErr)

	fmt.Println(wrappedErr)
	fmt.Println(unwrappedErr)

	// Output:
	// database failed: connection refused
	// connection refused
}

```

If the error does not wrap another error, `errors.Unwrap` returns `nil`.

```go
err := errors.New("plain error")
fmt.Println(errors.Unwrap(err)) // Output: <nil>

```

### Checking Wrapped Errors: `errors.Is`

Use `errors.Is` to check if an error (or any error wrapped inside it) matches a specific target error. This replaces direct `err == targetErr` comparisons.

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("resource not found")

func FetchData() error {
	// Simulating an error wrapped deep inside the application
	return fmt.Errorf("database layer: %w", ErrNotFound)
}

func main() {
	err := FetchData()

	// errors.Is digs through the wrapped layers to find ErrNotFound
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Handled appropriately: The resource wasn't found.")
	} else {
		fmt.Println("Some other error occurred.")
	}
}

```

### Extracting Custom Errors: `errors.As`

If you wrapped a *custom struct error* and need to access its specific fields later, use `errors.As`. This replaces manual type assertions (`err.(*MyType)`).

```go
package main

import (
	"errors"
	"fmt"
)

type NetworkError struct {
	StatusCode int
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error with status %d", e.StatusCode)
}

func makeRequest() error {
	return fmt.Errorf("API call failed: %w", &NetworkError{StatusCode: 503})
}

func main() {
	err := makeRequest()

	var netErr *NetworkError
	// errors.As looks for NetworkError in the chain and assigns it to netErr if found
	if errors.As(err, &netErr) {
		fmt.Printf("Caught a network error! Status Code: %d\n", netErr.StatusCode)
		// Output: Caught a network error! Status Code: 503
	}
}

```

---

## 5. Joining Multiple Errors (Go 1.20+)

Sometimes you run a batch of operations and want to collect *all* the errors that occurred instead of just returning the first one. Go provides `errors.Join` for this.

```go
package main

import (
	"errors"
	"fmt"
)

func validateUser() error {
	var errs error

	errs = errors.Join(errs, errors.New("email is invalid"))
	errs = errors.Join(errs, errors.New("password is too weak"))

	return errs
}

func main() {
	err := validateUser()
	if err != nil {
		fmt.Println("Validation failed:\n", err)
		// Output:
		// Validation failed:
		// email is invalid
		// password is too weak
	}
}

```

---

## `errors.Is` vs `errors.As`

Think of it this way:

* Use **`errors.Is`** when you are looking for a **specific instance** of an error (like a specific variable).
* Use **`errors.As`** when you are looking for a **specific type** of error (so you can cast it and read its unique fields).
* Use **`fmt.Errorf` with `%w`** when you want to **wrap** an error with extra context.
* Use **`errors.Unwrap`** when you want to manually get the **next error** inside a wrapped error.

---

## Simple `errors.Is` Example (Checking for a specific error)

Imagine you have a predefined global error (often called a sentinel error) for when an item is out of stock. You wrap it with extra context, but you still need to check if the root cause was that "out of stock" error.

```go
package main

import (
	"errors"
	"fmt"
)

// 1. Define a specific, reusable error variable
var ErrOutOfStock = errors.New("item is out of stock")

func buyItem() error {
	// 2. We wrap our specific error with more context using %w
	return fmt.Errorf("checkout failed: %w", ErrOutOfStock)
}

func main() {
	err := buyItem()

	// 3. Use errors.Is to see if ErrOutOfStock is anywhere inside that error chain
	if errors.Is(err, ErrOutOfStock) {
		fmt.Println("Show user: 'Sorry, this item sold out!'")
	} else {
		fmt.Println("Show user: 'Something went wrong with your payment.'")
	}
}

```

---

## Simple `errors.As` Example (Extracting a custom error type)

Imagine you created a custom error struct because you need to pass an HTTP status code along with the error. You want to extract that status code later.

```go
package main

import (
	"errors"
	"fmt"
)

// 1. Define a custom error type with a specific field (Code)
type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func fetchPage() error {
	// 2. Return the custom error, wrapped in extra context
	myErr := &HttpError{Code: 404, Message: "Page Not Found"}
	return fmt.Errorf("crashing during fetch: %w", myErr)
}

func main() {
	err := fetchPage()

	// 3. Create an empty pointer of your custom error type
	var targetErr *HttpError

	// 4. errors.As checks if the error matches the type, and if so,
	//    it automatically extracts it into 'targetErr'
	if errors.As(err, &targetErr) {
		fmt.Printf("Extracted successfully! HTTP Status Code was: %d\n", targetErr.Code)
	} else {
		fmt.Println("It was a completely different kind of error.")
	}
}

```

---

## Simple Wrap Example

Use `%w` inside `fmt.Errorf` to wrap an error.

```go
package main

import (
	"errors"
	"fmt"
)

func openFile() error {
	return errors.New("permission denied")
}

func loadUserProfile() error {
	err := openFile()
	if err != nil {
		return fmt.Errorf("load user profile failed: %w", err)
	}
	return nil
}

func main() {
	err := loadUserProfile()
	fmt.Println(err)
	// Output: load user profile failed: permission denied
}

```

---

## Simple `errors.Unwrap` Example

Use `errors.Unwrap` when you want to manually pull out the directly wrapped error.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	originalErr := errors.New("disk full")
	wrappedErr := fmt.Errorf("save report failed: %w", originalErr)

	fmt.Println("wrapped:", wrappedErr)
	fmt.Println("unwrapped:", errors.Unwrap(wrappedErr))

	// Output:
	// wrapped: save report failed: disk full
	// unwrapped: disk full
}

```

---

## All Together: `Is`, `As`, Wrap, and Unwrap

```go
package main

import (
	"errors"
	"fmt"
)

var ErrUnauthorized = errors.New("unauthorized")

type RequestError struct {
	StatusCode int
	Err        error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request failed with status %d: %v", e.StatusCode, e.Err)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

func callAPI() error {
	requestErr := &RequestError{
		StatusCode: 401,
		Err:        ErrUnauthorized,
	}

	// Wrap the custom error with more context.
	return fmt.Errorf("call API: %w", requestErr)
}

func main() {
	err := callAPI()

	// 1. Wrap: callAPI returned an error wrapped with fmt.Errorf and %w.
	fmt.Println(err)

	// 2. Is: check whether ErrUnauthorized exists anywhere in the chain.
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("errors.Is: user must log in again")
	}

	// 3. As: extract a specific error type from anywhere in the chain.
	var requestErr *RequestError
	if errors.As(err, &requestErr) {
		fmt.Println("errors.As: status code is", requestErr.StatusCode)
	}

	// 4. Unwrap: get the next directly wrapped error.
	nextErr := errors.Unwrap(err)
	fmt.Println("errors.Unwrap:", nextErr)
}

```
