## Error Handling in Go

**Rules for Error Handling in Go:**
1. Error Stops here
2. Going to log it
3. Make a decision around recovery or Shutdown
If some piece of code not going to handle the error, it should `Wrap` all the context it and return it to the caller.
