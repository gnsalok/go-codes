# Go Concurrency: Buffered vs. Unbuffered Channels

## 1. Unbuffered Channels (`make(chan T)`)

**The Concept:**  
An unbuffered channel has no capacity to store values. It requires both a sender and a receiver to be ready at the exact same instant. This is known as a Rendezvous.

**Analogy: A Phone Call.**  
You cannot speak (send) unless someone is on the other line listening (receive). If you talk and no one is there, you wait (block) until they pick up.

**When to use:**
- Strict Synchronization: When you need to guarantee that a specific step is finished before moving on.
- Hand-off: When you want to ensure the ownership of data is passed completely from one goroutine to another.
- Default Choice: Always start with unbuffered channels to avoid complexity. Only add buffers if you have a specific performance reason.

## 2. Buffered Channels (`make(chan T, size)`)

**The Concept:**  
A buffered channel has an internal queue. The sender can send values without blocking, as long as the buffer isn't full. The receiver can receive values without blocking, as long as the buffer isn't empty.

**Analogy: Sending an Email.**  
You can click send (send) even if the recipient is asleep. The email sits in their inbox (buffer). You only get blocked if their inbox is completely full (storage quota reached).

**When to use:**
- Decoupling: To separate the speed of the producer from the speed of the consumer.
- Burst Handling: When requests come in spikes, but processing is steady. The buffer absorbs the spike.
- Limiting Concurrency (Semaphores): Using a buffered channel of size N to ensure only N goroutines run at once.

**Warning for Interviews:**  
Do not use buffered channels just to "fix" a deadlock. If your logic is broken, a buffer usually just delays the deadlock or creates a race condition.

## 3. How to combine Goroutines and Channels

The most common pattern for Senior Engineers to know is the Fan-Out / Fan-In (Worker Pool).

- The Job Queue (Buffered Channel): Holds the tasks to be done.
- The Workers (Goroutines): A fixed number of goroutines reading from the Job Queue.
- The Results (Buffered Channel): Where workers send their output.
- The Manager (WaitGroup): Ensures all workers finish before closing the results.
