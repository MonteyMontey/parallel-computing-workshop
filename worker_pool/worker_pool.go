package worker_pool

import (
	"fmt"
	"time"
)

// the direction of the arrow is just a constraint how the channel is used
// <- chan = reading from channel
// chan <- = writing to channel
// chan = both reading and writing

func worker(id int, jobs <-chan int, results chan<- int) {

	// iteration doesn't stop until job channel is closed
	for j := range jobs {
		fmt.Println(j)
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func RunWorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// creating the worker pool
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// creating jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// printing results
	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}

// runtime already maintains an internal thread pool of goroutines

// you can use a semaphore to constraints the amount of goroutines currently used
