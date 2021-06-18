package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func worker1(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker Id -- ", id, " Job started ", j)
		r := calc(id, j)
		wg.Done()
		time.Sleep(time.Second)
		fmt.Println("Worker Id -- ", id, " Job Finished ", j)
		results <- r
	}
}

func calc(opr int, data int) int {
	switch opr {
	case 1:
		return data + data
	case 2:
		return data - data
	case 3:
		return data * data
	case 4:
		return data / data
	}
	return 0
}

func DoWeightGroup() {
	fmt.Println("-------------DoWeightGroup----------------")
	jobs := make(chan int, 5)
	result := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		go worker1(i, &wg, jobs, result)

	}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 5; i++ {
		r := <-result
		fmt.Println("Result ", r)
	}
	wg.Wait()
}
