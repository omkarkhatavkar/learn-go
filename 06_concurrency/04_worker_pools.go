package main

import (
	"fmt"
	"time"
)

// A worker pool is a design pattern that uses a fixed number of worker goroutines
// to process a stream of jobs from a channel.

// A worker goroutine. It listens for jobs on the `jobs` channel and sends the result
// to the `results` channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Simulate work by sleeping.
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send the result to the results channel.
	}
}

func main() {

	// --- Worker Pool Example ---
	// We'll create a job queue and a results channel.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 worker goroutines.
	// These workers are ready to receive jobs from the `jobs` channel.
	numWorkers := 3
	fmt.Printf("Starting %d workers...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send 9 jobs to the `jobs` channel.
	// The jobs are put into the channel's buffer and can be picked up by any worker.
	numJobs := 9
	fmt.Printf("\nSending %d jobs...\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Close the `jobs` channel. This signals that no more jobs will be sent.
	// Without this, the `for job := range jobs` loop in the worker will never terminate.
	close(jobs)

	// Collect all the results from the `results` channel.
	fmt.Println("\nCollecting results...")
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("Collected result: %d\n", result)
	}
	fmt.Println("All results collected.")
}
