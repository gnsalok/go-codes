package main

import (
	"fmt"
)

func openFile(fileName string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error opening file:", err)
		}
	}()

	// Simulate an error opening the file
	if fileName == "invalid.txt" {
		panic("Unable to open file")
	}

	fmt.Println("File", fileName, "opened successfully")
	return nil
}

func main() {
	err := openFile("valid.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = openFile("invalid.txt") // This will trigger the panic and recovery
	// Output: Error opening file: Unable to open file
}

/*
In the provided code, the `recover` function is used to handle a **panic** gracefully. Here's a step-by-step explanation of what is happening:

### What is `recover`?
`recover` is a built-in Go function that allows you to regain control of a program after a panic occurs. A **panic** is a runtime error that causes the program to stop executing normally. When `recover` is called inside a `defer` function, it captures the value passed to `panic` and prevents the program from crashing.

### How `recover` works in your code:
1. **Deferred Function**:
   ```go
   defer func() {
       if err := recover(); err != nil {
           fmt.Println("Error opening file:", err)
       }
   }()
   ```
   - The `defer` keyword ensures that the anonymous function is executed at the end of the `openFile` function, regardless of whether it exits normally or due to a panic.
   - Inside this deferred function, `recover` is called to check if a panic occurred.

2. **Capturing the Panic**:
   - If a panic occurs (e.g., `panic("Unable to open file")`), `recover` will capture the panic value (in this case, the string `"Unable to open file"`).
   - The captured value is assigned to `err`, and the program continues executing the deferred function instead of crashing.

3. **Handling the Panic**:
   - The deferred function prints an error message using the captured panic value:
     ```go
     fmt.Println("Error opening file:", err)
     ```

4. **Normal Execution**:
   - If no panic occurs, `recover` returns `nil`, and the deferred function does nothing.

### Example Walkthrough:
- If you call `openFile("valid.txt")`:
  - No panic occurs.
  - The deferred function runs, but `recover` returns `nil`, so nothing is printed.
  - The function completes normally.

- If you call `openFile("invalid.txt")`:
  - A panic is triggered with the message `"Unable to open file"`.
  - The deferred function runs, `recover` captures the panic value, and the error message is printed:
    ```
    Error opening file: Unable to open file
    ```
  - The program does not crash and continues executing.

### Why use `recover` here?
The `recover` function is used to ensure that the program can handle unexpected errors (like trying to open an invalid file) without crashing entirely. It provides a way to gracefully recover from panics and continue execution.

### Gotcha:
- `recover` only works if it is called within a deferred function. If you call it outside of `defer`, it will always return `nil`.  - The program does not crash and continues executing.

### Why use `recover` here?
The `recover` function is used to ensure that the program can handle unexpected errors (like trying to open an invalid file) without crashing entirely. It provides a way to gracefully recover from panics and continue execution.

### Gotcha:
- `recover` only works if it is called within a deferred function. If you call it outside of `defer`, it will always return `nil`.
*/
