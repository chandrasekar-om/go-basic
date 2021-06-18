package concurrency

import (
	"fmt"
	"time"
)

func DoChannels() {
	message := make(chan string)

	go func(argv string) {
		message <- argv
	}("Hello")

	msg := <-message
	fmt.Println(msg)
}

func DoChannelBuffering() {
	message := make(chan string, 4)
	message <- "Harish"
	message <- "Harini"
	message <- "Srinikethan"
	message <- "Srikrithi"

	for i := 0; i < 4; i++ {
		fmt.Println(<-message)
	}
}

func add(done chan bool, argv int) {
	total := 0
	for i := 1; i < argv; i++ {
		total += i
	}
	fmt.Println("add process ", total)
	done <- true
}

func sub(done chan bool, argv int) {
	total := 0
	for i := 1; i < argv; i++ {
		total -= i
	}
	fmt.Println("sub process ", total)
	done <- true
}

func DoChannelSynchronization() {
	done := make(chan bool)

	go add(done, 5)
	go sub(done, 5)

	time.Sleep(time.Second)
	<-done
}
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func DoChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello pingpong")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func DoSelectWithChannels() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(msg string) {
		time.Sleep(1 * time.Second)
		chan1 <- msg
	}("message1")
	go func(msg string) {
		time.Sleep(2 * time.Second)
		chan2 <- msg
	}("message2")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}

}

func DoTimeoutsWithChannels() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go func(msg string) {
		time.Sleep(1 * time.Second)
		chan1 <- msg
	}("message1")
	go func(msg string) {
		time.Sleep(3 * time.Second)
		chan2 <- msg
	}("message2")
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout....")
		}
	}
}

func DoNonBlockingChannel() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

func DoClosingChannels() {
	tasks := make(chan int, 5)
	signal := make(chan bool)

	go func() {
		for {
			t, more := <-tasks
			if more {
				fmt.Println("We got task ", t)
			} else {
				fmt.Println("All the task completed")
				signal <- true
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		tasks <- i
		fmt.Println("Send the job ", i)
	}
	close(tasks)

	<-signal
}

func DoRangeoverChannels() {
	message := make(chan string, 3)
	message <- "one"
	message <- "two"
	message <- "three"
	close(message)

	for v := range message {
		fmt.Println(v)
	}
}
