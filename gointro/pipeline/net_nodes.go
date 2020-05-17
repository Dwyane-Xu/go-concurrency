package pipeline

import (
	"bufio"
	"net"
)

// NetworkSink 传入端口号和一个channel，监听端口号开启tco连接
// 当有客户端连接时，将channel中的数据写到连接中
func NetworkSink(addr string, in <-chan int) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()

		WriterSink(writer, in)
	}()
}

// NetworkSource 传入端口号，通过端口号连接服务端
// 获取服务端传过来的数据到channel中，返回channel
func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()

	return out
}
