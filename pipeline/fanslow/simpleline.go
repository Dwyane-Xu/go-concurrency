package main

import (
	"fmt"
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
			// time.Sleep(time.Second)
		}
	}()

	return out
}

// go run simpleline.go  8.69s user 3.26s system 169% cpu 7.032 total
func main() {
	in := producer(10000000)
	ch := square(in)

	// consumer
	for range ch {
		// fmt.Printf("%3d", ret)
	}
	fmt.Println()
}
