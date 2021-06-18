package timers

import (
	"fmt"
	"time"
)

func DoTimers() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	fmt.Println("Timer 1 is fired")

	timer2 := time.NewTimer(time.Second)

	go func(st string) {
		<-timer2.C
		fmt.Println(st)
	}(" Time 2 is fired ")
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 is stoped")
	}
	time.Sleep(time.Second)
}

func DoTickers() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Ticker is ticking ", t)
			}
		}
	}()
	time.Sleep(time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker is stoped")
}
