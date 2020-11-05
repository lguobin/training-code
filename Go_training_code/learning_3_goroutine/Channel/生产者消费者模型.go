package Channel

import (
	"fmt"
	"time"
)

// 生产者 - 只写管道
func producer(ch chan<- int) {
	i := 1
	for {
		ch <- i
		fmt.Println("生产者 -> Send:", i)
		i++
		time.Sleep(time.Second + 1) // 避免数据流动过快
	}
}

// 消费者 - 只读管道
func consumer(ch <-chan int) {
	for {
		v := <-ch
		fmt.Println("消费者 -> Receive: ", v)
		time.Sleep(time.Second + 2)
	}
}

func Chan_Demo_6() {
	fmt.Println("生产者消费者模型")
	ch := make(chan int, 5)
	go producer(ch)
	go consumer(ch)
	for {
	}
}
