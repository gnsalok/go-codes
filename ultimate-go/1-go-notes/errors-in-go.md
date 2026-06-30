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

## 4. Modern Error Handling: Wrapping and Unwrapping (Go 1.13+)

When an error passes up through multiple layers of your application, you often want to add context to it without losing the original error. Go allows you to **wrap** errors using `%w` in `fmt.Errorf`.

### Wrapping Errors

```go
originalErr := errors.New("permission denied")
// Wrapping the original error with more context
wrappedErr := fmt.Errorf("failed to delete file: %w", originalErr)

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

## Summary Best Practices

| What to do | How to do it |
| --- | --- |
| **Simple text error** | Use `errors.New("message")` |
| **Dynamic text error** | Use `fmt.Errorf("failed: %s", detail)` |
| **Add context to an existing error** | Use `fmt.Errorf("context: %w", err)` |
| **Check if an error matches a sentinel** | Use `errors.Is(err, target)` |
| **Extract a custom error type** | Use `errors.As(err, &customErrTarget)` |

Which specific part of error handling—like writing clean code to avoid "if err != nil" boilerplate, or custom error architectures—would you like to dive into next?