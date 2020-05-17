package main

import (
	"bufio"
	"fmt"
	"os"

	"go-concurrency/gointro/pipeline"
)

func main() {
	file, err := os.Create("./gointro/large.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 产生指定数量的随机数，返回channel
	p := pipeline.RandomSource(100000000)

	// 传入channel，将产生的随机数保存到文件中
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	// 打开文件
	file, err = os.Open("./gointro/large.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件中的数据，返回channel
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)

	// 遍历channel，打印100个数
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(1, 5, 3, 9, 2, 15, 4)))
	for v := range p {
		fmt.Println(v)
	}
}
