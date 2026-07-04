## Channels 

Table of Contents:
- [Channels](#channels)
- [Channel Semantics](#channel-semantics)
- [Fan-out pattern with buffer channel](#fan-out-pattern-with-buffer-channel)
- [Important : Channels and WaitGroups solve different problems](#important--channels-and-waitgroups-solve-different-problems)
- [When to think about using Channels](#when-to-think-about-using-channels)


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


---

# When to think about using Channels

It is perfectly normal to feel this way. Most developers come from a "sequential" background where we think in terms of **calling functions** (Action A -> Result -> Action B). In concurrency, you have to think in terms of **passing batons**.

To help you build a mental "trigger" for channels, here is a 4-point checklist. If your problem has any of these, a channel is likely the answer.

---

## 1. The "Pipeline" Trigger (Hand-offs)

**Ask yourself:** "Does data move through stages?"
If your logic looks like: *Fetch Data* -> *Transform It* -> *Save to DB*, and you want these to happen at the same time on different items.

* **Mental Image:** An assembly line.
* **Why Channels?** Stage 1 sends the "processed" item to Stage 2 via a channel. Stage 1 doesn't care what Stage 2 does; it just moves to the next item.

## 2. The "Result Collector" Trigger

**Ask yourself:** "Am I starting 10 workers and I need their answers back?"
If you use a `sync.WaitGroup`, you know when they are *done*, but you can't easily get their data back safely without a Mutex.

* **Mental Image:** 10 researchers going out to find info and coming back to drop their reports on one desk.
* **Why Channels?** Each worker sends its result into a single `resultsChan`. The main goroutine just sits and reads from that channel until it's empty.

## 3. The "Orchestration" Trigger (Signal Events)

**Ask yourself:** "Does one part of my code need to tell another part to 'Start' or 'Stop'?"

* **Mental Image:** A starter pistol at a race.
* **Why Channels?** An empty `chan struct{}` is the cheapest way to send a signal. Closing a channel is a "broadcast" signal that tells every listening goroutine: "The job is over, clean up and exit."

## 4. The "Throttler" Trigger (Scarcity)

**Ask yourself:** "Do I need to limit how many things happen at once?"

* **Mental Image:** A nightclub with a bouncer at the door.
* **Why Channels?** A **buffered channel** acts as the bouncer. If the buffer is 5, only 5 goroutines can "get in" (send to the channel). The 6th must wait until someone "leaves" (receives from the channel).

---

## The "Checklist" Decision Matrix

| If you need to... | Use a Channel (`chan`) | Use a Mutex (`sync.Mutex`) |
| --- | --- | --- |
| **Pass data ownership** | ✅ Yes | ❌ No |
| **Coordinate a sequence** | ✅ Yes | ❌ No |
| **Distribute work** | ✅ Yes | ❌ No |
| **Protect a simple counter** | ❌ No (Too slow) | ✅ Yes |
| **Protect a Map/Cache** | ❌ No | ✅ Yes |
| **High-perf state management** | ❌ No | ✅ Yes |

---

## A Simple Mental Exercise

Next time you write a `for` loop that does something slow (like an API call), stop and ask:

1. **Can I run this in a goroutine?** (Usually yes).
2. **How do I get the result back?** (If the answer is "I'll append it to a slice," **STOP**. You can't append to a slice safely from multiple goroutines. **This is your trigger to use a channel instead.**)

### Example: The "Wait & Collect" Pattern

This is the most common way to start thinking in channels:

```go
func fetchAll(urls []string) []string {
    results := make(chan string, len(urls)) // Trigger 2: Result Collector

    for _, url := range urls {
        go func(u string) {
            res := download(u)
            results <- res // Handing the baton back
        }(url)
    }

    // Collect them
    var final []string
    for i := 0; i < len(urls); i++ {
        final = append(final, <-results) // Pulling from the "desk"
    }
    return final
}

```