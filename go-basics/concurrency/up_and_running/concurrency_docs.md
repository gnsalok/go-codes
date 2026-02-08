## Concurrency 

The composition of independently executing processes. - Rob Pike


## Trends 

Why concurrency is important? see the below trends, no. of logical cores in computer started increase and single thread performance and frequency keep decreasing.

![alt text](image.png)

## WaitGroup 

Golang Waitgroup allows you to block a specific code block to allow a set of goroutines to complete execution. 
it is used for synchronization.


## Channel 

**"Don't communicate by sharing memory, share memory by communicating." - Rob Pike Go Proverb**

Channels are used to pass data between / synchronize goroutine, including func main.

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.


```go 

ch := make(chan string)

ch <- mydata    // BLOCKING 

myVar <- ch  // BLOCKING

close(ch)   // (optional)

```

> Note: You can monitor multiple channel using `select` statement.




