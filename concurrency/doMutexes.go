package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func DoMutexes() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for i := 0; i < 100; i++ {

		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	finalReadOps := atomic.LoadUint64(&readOps)
	fmt.Println("Final Read ", finalReadOps)
	finalWriteOps := atomic.LoadUint64(&writeOps)
	fmt.Println("Final Write ", finalWriteOps)
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}
