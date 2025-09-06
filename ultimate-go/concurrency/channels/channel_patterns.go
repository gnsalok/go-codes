// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// ------- `waitForTask` pattern can be used for Pooling
	// waitForTask()
	// pooling()

	// ------- `waitForResult` pattern can be used for "Drop" and "Fan-out" pattern. ****
	// waitForResult()
	// fanOut()
	drop()

	// More advanced patterns
	// fanOutSem()
	// boundedWorkPooling()

	// waitForFinished()
	// Cancellation Pattern
	//cancellation()

	// ------- Retry Pattern
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// retryTimeout(ctx, time.Second, func(ctx context.Context) error { return errors.New("always fail") })

	// ------- Channel Cancellation
	//	stop := make(chan struct{})
	// channelCancellation(stop)
}

// waitForTask: In this pattern, the parent goroutine sends a signal to a
// child goroutine for waiting to be told what to do.
func waitForTask() {
	ch := make(chan string)

	go func() {
		data := <-ch // RECEIVE :: this is wait for task ; add unknown latency. Guarantees.
		fmt.Println("child : recv'd signal :", data)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data" // SEND
	fmt.Println("parent : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// waitForResult: In this pattern, the parent goroutine waits for the child
// goroutine to finish some work to signal the result.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// Note : you can better achieve this with waitGroup
func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		close(ch)
		fmt.Println("employee : sent signal")
	}()

	_, wd := <-ch // if signal with data then wd will be true; this is boolean flag, no data to receive.
	fmt.Println("manager : received signal : ", wd)

	time.Sleep(1 * time.Second)
}

// pooling: In this pattern, the parent goroutine signals 100 pieces of work
// to a pool of child goroutines waiting for work to perform.
// This is kind of WaitForTask Patterns
func pooling() {
	ch := make(chan string)

	const emps = 2

	for e := 0; e < emps; e++ {
		// launches 2 goroutine
		go func(emp int) {
			// This will keep polling to get task ; and get closed when channel is closed.
			// getting data from channel :: putting channel in wait state to get Task
			for p := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// fanOut: In this pattern, the parent goroutine creates 2000 child goroutines
// and waits for them to signal their results.
// This is kind of WaitForResult Patterns
func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for c := 0; c < emps; c++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", emp)
		}(c)
	}

	for emps > 0 {
		d := <-ch // receive
		emps--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func fanOutSem() {
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			<-sem
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// boundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signalling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// drop: In this pattern, the parent goroutine signals 2000 pieces of work to
// a single child goroutine that can't handle all the work. If the parent
// performs a send and the child is not ready, that work is discarded and dropped.

// Rate limit you can assume here
func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	const work = 300

	// for select allow to do multiple channel operation
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default: // default case will execute when we are unable to send data into buffered channel and drop the request
			fmt.Println("parent : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// cancellation: In this pattern, the parent goroutine creates a child
// goroutine to perform some work. The parent goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the parent goroutine walks away.
func cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// retryTimeout: You need to validate if something can be done with no error
// but it may take time before this is true. You set a retry interval to create
// a delay before you retry the call and you use the context to set a timeout.
func retryTimeout(ctx context.Context, retryInterval time.Duration, check func(ctx context.Context) error) {

	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}

		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}

		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)

		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

// channelCancellation shows how you can take an existing channel being
// used for cancellation and convert that into using a context where
// a context is needed.
func channelCancellation(stop <-chan struct{}) {

	// Create a cancel context for handling the stop signal.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// If a signal is received on the stop channel, cancel the
	// context. This will propagate the cancel into the p.Run
	// function below.
	go func() {
		select {
		case <-stop:
			cancel()
		case <-ctx.Done():
		}
	}()

	// Imagine a function that is performing an I/O operation that is
	// cancellable.
	func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.ardanlabs.com/blog/index.xml", nil)
		if err != nil {
			return err
		}
		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		return nil
	}(ctx)
}
