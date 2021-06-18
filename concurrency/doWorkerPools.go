package concurrency

import (
	"fmt"
	"time"
)

func DoWorkerPools() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= 5; i++ {
		r := <-results
		fmt.Println("Result ", r)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker Id ", id, " Job started ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker Id ", id, " Job Finished ", j)
		results <- j * 2
	}
}
