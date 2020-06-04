package main

import "fmt"

func main() {
	ch1 := gen()
	ch2 := gen()

	out := combine(ch1, ch2)

	for x := range out {
		fmt.Println(x)
	}
}

func gen() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 3; i++ {
			x := i
			ch <- x
		}
	}()

	return ch
}

func combine(inCh1, inCh2 <-chan int) <-chan int {
	// 输出通道
	out := make(chan int)

	// 启动协程合并数据
	go func() {
		defer fmt.Println("combine exit")
		defer close(out)

		for {
			select {
			case x, open := <-inCh1:
				if !open {
					inCh1 = nil
					break
				}
				out <- x
			case x, open := <-inCh2:
				if !open {
					inCh2 = nil
					break
				}
				out <- x
			}

			// 当ch1和ch2都关闭时才退出
			if inCh1 == nil && inCh2 == nil {
				break
			}
		}
	}()

	return out
}
