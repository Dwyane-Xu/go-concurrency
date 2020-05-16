package main

import (
	"sync"
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

func merge(bufSize int, cs ...<-chan int) <-chan int {
	out := make(chan int, bufSize)
	var wg sync.WaitGroup

	// FAN-IN
	for _, c := range cs {
		wg.Add(1)
		go func(in <- chan int) {
			defer wg.Done()
			for n := range in {
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Fanline(bufSize int) {
	in := producer(10000)

	// FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for range merge(bufSize, c1, c2, c3) {
	}
}

func SimpleLine() {
	in := producer(10000)
	ch := square(in)

	// consumer
	for range ch {
	}
}



