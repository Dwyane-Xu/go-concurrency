package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
			// simulate
			time.Sleep(time.Second)
		}
	}()

	return out
}

// go run simpleline.go  0.23s user 0.19s system 3% cpu 12.407 total
func main() {
	in := producer(12)
	ch := square(in)

	// consumer
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()
}
