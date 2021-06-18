package concurrency

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func DoStatefulGoroutines() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Second)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Second)
			}
		}()
	}
	time.Sleep(time.Second)
	readR := atomic.LoadUint64(&readOps)
	writeW := atomic.LoadUint64(&writeOps)
	fmt.Println("Read Final ", readR)
	fmt.Println("Write Final ", writeW)
}
