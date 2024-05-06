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
In Golang, the `recover` function is used for error handling in conjunction with the `panic` and `defer` statements. It allows you to regain control of a program that has encountered a panic and potentially handle the error gracefully.

Here's a breakdown of how it works:

**Panic:**

- The `panic` statement is used to signal a critical error condition that should stop the normal execution of the program.
- When a panic occurs, the program starts unwinding the call stack, cleaning up resources as it goes.

**Recover:**

- The `recover` function is a built-in function that attempts to capture the value passed to the `panic` statement from within a `defer` function.
- A `defer` statement specifies a function to be called after the surrounding function finishes executing, regardless of how it exits (normally, by returning, or abnormally, by panicking).

**Putting It Together:**

1. A function panics with a specific error message.
2. A `defer` statement is used to schedule the execution of another function after the main function finishes or panics.
3. Inside the `defer` function, you can call `recover`.
4. If a panic occurred before reaching the `defer` statement, `recover` will capture the error message passed to the `panic` and return it.
5. You can then use the recovered value to handle the error gracefully, potentially cleaning up resources or logging the issue.


In this example:

1. The `openFile` function attempts to open a file.
2. It uses a `defer` statement to schedule a cleanup function.
3. Inside the cleanup function, `recover` is used to capture any potential panic.
4. If `openFile` panics due to the invalid file name, `recover` captures the error message ("Unable to open file") within the cleanup function.
5. The error message is then printed.

**Important Points:**

- `recover` only works within a `defer` function. Calling it outside of `defer` will not capture a panic.
- `recover` only retrieves the value from the most recent panic. If there are nested panics, only the value from the outermost panic is accessible.
- Using `recover` doesn't prevent the program from terminating after a panic. However, it allows you to potentially perform cleanup actions before termination.

Remember, `panic` and `recover` are generally not the preferred way to handle errors in Go. It's recommended to use Go's built-in error handling mechanisms for most cases. The `recover` function is typically used for exceptional situations where you need to handle unexpected errors or system issues.

*/
