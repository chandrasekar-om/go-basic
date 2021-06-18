package concurrency

import (
	"fmt"
	"time"
)

func greet(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": Hello :", i)
	}
}

func DoGoroutine() {
	greet("Direct")

	go greet("Goroutine")

	go func(argv string) {
		fmt.Println(argv)
	}("anonymous")

	time.Sleep(time.Second)
	fmt.Println("Done!")
}
