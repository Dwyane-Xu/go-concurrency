package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Second)
		}
	}()
	return out
}

// consumer only read data from in channel and print it
func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})
	t := time.Tick(time.Millisecond * 500)
	processedCnt := 0

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		// in for-select using ok to exit goroutine
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Process %d\n", x)
				processedCnt++
			case <-t:
				fmt.Printf("Working, processedCnt = %d\n", processedCnt)
			}
		}
	}()

	return finish
}

func main() {
	out := producer(3)
	finish := consumer(out)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}

// go run exit.go
// send 0
// Process 0
// send 1
// Process 1
// send 2
// Process 2
// producer exit
// worker exit
// main exit
