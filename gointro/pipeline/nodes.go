package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

// Init 初始化，保存执行的时间
func Init() {
	startTime = time.Now()
}

// ArraySource 开启一个协程：读取输入的数据到channel中，返回channel
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// InMemSort 开启一个协程：从channel中读取数据，对数据进行排序
// 排序后的数据传到一个新的channel中，返回channel
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		var a []int
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))

		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))

		for v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// MergeN 传入多个channel，递归调用，进行归并排序
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}

// Merge 开启一个协程：从两个channel中读取数据
// 把数据合并到一个新的channel中，返回channel
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()

	return out
}

// ReaderSource 开启一个协程：从文件中读取数据到channel中，返回channel
// chunkSize 是文件块大小
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

// WriterSink 从channel中读取数据，并写到文件中
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// RandomSource 开启一个协程：生成count个随机数，并发送到一个channel中，返回channel
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}
