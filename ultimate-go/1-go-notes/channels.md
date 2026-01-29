## Channels 

- Channels is how you deal with orchestration in Go. [Read more about channels](https://github.com/gnsalok/gotraining/blob/master/topics/go/concurrency/channels/README.md)

- Channels allow goroutines to communicate with each other through the use of `signaling semantics`. 
- Channels accomplish this signaling through the use of sending/receiving data or by identifying state changes on individual channels. 
- Don't architect software with the idea of channels being a queue, focus on signaling and the semantics that simplify the orchestration required.


### Channel Semantics

- When you think about Channels always think about Signaling is the semantic.
    - Pointers is for sharing and Channels are for Signaling.

- We can signal with Data or without data.
- We also need to understand Guarantee?
    - Does the goroutine sending a signal need a guarantee that signal's been received? Use Unbuffered channel. But we have live with Unknown latency.
    - But we can reduce those latency but without guarantees (Buffer channels). 
        - Buffers don't provide performance (Big buffer). But small buffers reduce the latency.





``` go 
<-ch // receive
ch <- "data" // send

```


### Fan-out pattern with buffer channel 

> See code in `concurrency/channels/channel_patterns.go`

- Channel buffer should be a reasonable number. It should not any random no.
- To reduce the unknown latency between send and receiving we use buffer channel. Make sure sending to channel is not block.
- If anytime a buffer channel get full we've ask ourselves what happens when thi signaling I am gonna send blocks?
- It should be like : 1 : 1 ; One buffer to one goroutine.

---

## Important : Channels and WaitGroups solve different problems:

- Use sync.WaitGroup when you only need to wait for a set of goroutines to finish. It’s about completion, not communication.

- Use channels when you need communication, coordination, or data flow between goroutines **(fan-out/fan-in, pipelines, cancellation, backpressure, streaming results)**.
Concrete example: collecting results as they finish (not just waiting):

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, out chan<- int) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	out <- id * id
}

func main() {
	rand.Seed(time.Now().UnixNano())

	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		go worker(i, results)
	}

	// Channel lets us receive results as they complete.
	for i := 0; i < 3; i++ {
		fmt.Println("got", <-results)
	}
}
```

If you used a WaitGroup here, you’d still need a channel (or shared memory + locks) to get the results. So the channel is the right primitive because communication is the goal.

Rule of thumb: **WaitGroup = “wait for done”**; **Channel = “share/stream data or coordinate.”**