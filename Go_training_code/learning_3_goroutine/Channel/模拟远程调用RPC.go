package Channel

import (
	"fmt"
)

/**
模拟远程调用RPC：使用通道代替 Socket 实现 RPC 的过程。
客户端与服务器运行在同 一个进程， 服务器和客户端在两个 goroutine 中运行。
*/

// 模拟客户端
func RPCClient(ch chan string, req string) (string, error) {
	ch <- req // 向服务器发送请求模拟：请求数据放入通道
	select {  // 等待服务器返回模拟：使用select
	case data := <-ch:
		return data, nil
	}
}

// 模拟服务端
func RPCServer(ch chan string) {
	// 通过无限循环继续处理下一个客户端请求。
	for {
		data := <-ch
		fmt.Println("server received: ", data)
		ch <- "roger" // 向客户端反馈
	}
}

func Chan_Demo_8() {

	// 模拟 启动服务器
	ch := make(chan string)
	go RPCServer(ch)

	// 模拟 发送数据
	receive, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client receive: ", receive)
	}
}
