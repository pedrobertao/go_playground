package main

import (
	"fmt"
)

// Result holds the result of a work function execution
type ResultPromise struct {
	Result any
	Error  error
}

// Work is a function type that returns a result and an error
type Work func() (any, error)

// Promise is a worker goroutine that executes work from a jobs channel
// and sends the result to a results channel
func Promise(id int64, jobs <-chan Work, results chan<- ResultPromise) {
	fmt.Printf("[%d] Fechting...\n", id)
	// Get work from the jobs channel
	w := <-jobs
	// Execute the work function
	exp, err := w()
	// Send the result to the results channel
	results <- ResultPromise{
		Result: exp,
		Error:  err,
	}
}

// PromiseAll executes all work functions concurrently and returns all results
func PromiseAll(works []Work) []ResultPromise {
	// Create buffered channels for jobs and results
	jobs := make(chan Work, len(works))
	results := make(chan ResultPromise, len(works))

	// Start a worker goroutine for each work function
	for i := range works {
		go Promise(int64(i+1), jobs, results)
	}

	// Send all work functions to the jobs channel
	for i := 0; i < len(works); i++ {
		jobs <- works[i]
	}

	// Close the jobs channel to signal no more work
	close(jobs)

	// Collect all results from the workers
	res := make([]ResultPromise, len(works))
	for i := 0; i < len(works); i++ {
		workDone := <-results
		res = append(res, workDone)
	}
	return res
}

func ExamplePromiseAll() {
	// Execute HTTP GET requests concurrently using the Promise pattern
	results := PromiseAll([]Work{
		func() (any, error) {
			return httpGet("https://www.google.com")
		},
		func() (any, error) {
			return httpGet("https://www.amazon.com")
		},
		func() (any, error) {
			return httpGet("https://www.facebook.com")
		},
	})

	// Print the results
	fmt.Println(">", results)
}
