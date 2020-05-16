package main

import (
	"fmt"
	"time"
)

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")

		t := time.Tick(time.Millisecond * 500)

		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-t:
				fmt.Println("Working .")
			}
		}
	}()
	return
}

func main() {
	stopCh := make(chan struct{})
	worker(stopCh)

	time.Sleep(time.Second * 2)
	close(stopCh)

	// Wait some print
	time.Sleep(time.Second)
	fmt.Println("main exit")
}

// go run exit.go
// Working .
// Working .
// Working .
// Working .
// Recv stop signal
// worker exit
// main exit
