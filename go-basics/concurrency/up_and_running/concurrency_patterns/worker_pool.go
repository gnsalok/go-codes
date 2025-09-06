// Worker Pool is one of the concurrency pattern

package main

import (
	"fmt"
	"time"
)

func main() {
	const noOfJobs = 10
	jobsChan := make(chan int, noOfJobs)

	completedJobChan := make(chan int, noOfJobs)

	for w := 1; w <= 3; w++ { // No. of worker is 3, and each worker is a goroutine
		go worker(w, jobsChan, completedJobChan)

	}

	for j := 1; j <= noOfJobs; j++ {
		jobsChan <- j // this loads the job into the jobs channel
	}

	close(jobsChan) // when all jobs are loaded close the channel,  meaning nothing can be loaded into the channel.

	for a := 1; a <= noOfJobs; a++ {
		<-completedJobChan // Read the completed job from the channel and do nothing with the context. Close of the program.

	}
}

func worker(id int, jobsChan <-chan int, completedJobChan chan<- int) {

	for j := range jobsChan {
		fmt.Println("worker ", id, " started job : ", j, " wih ", len(jobsChan), " job left to the process")
		time.Sleep(2 * time.Second)
		fmt.Println("worker ", id, " finished the job ", j)
		completedJobChan <- j
	}

}
