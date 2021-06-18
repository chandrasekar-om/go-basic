package concurrency

import (
	"fmt"
	"time"
)

func DoRateLimiting() {
	doBasicRateLimiting()
	doShortBursts()
}

func doBasicRateLimiting() {
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)
	timer := time.Tick(200 * time.Millisecond)

	for r := range request {
		<-timer
		fmt.Println("Request ", r, time.Now())
	}
}

func doShortBursts() {
	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for r := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- r
		}
	}()
	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for r := range burstyRequest {
		<-burstyLimiter
		fmt.Println("Bursty Request ", r, time.Now())
	}
}
