package Channel

import (
	"fmt"
	"time"
)

// 写端
func write(ch chan int) {
	ch <- 100
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 200
	fmt.Printf("ch addr：%v\n", ch) // 输出内存地址
	ch <- 300                      // 该处数据未读取，后续操作直接阻塞
	fmt.Printf("ch addr：%v\n", ch) // 没有输出
}

// 读端
func read(ch chan int) {
	// 只读取两个数据
	fmt.Printf("取出的数据data1：%v\n", <-ch) // 100
	fmt.Printf("取出的数据data2：%v\n", <-ch) // 200
}
func TestChannel() {
	ch := make(chan int)
	// 向协程中写入数据
	go write(ch)
	// 向协程中读取数据
	go read(ch)
	time.Sleep(time.Millisecond * 100)

	N_ch := make(chan int, 1)
	// 向协程中写入数据
	go write(N_ch)
	time.Sleep(time.Millisecond * 100)
}
func TestChannel_Status() {
	ch := make(chan int, 10)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	for {
		if count, ok := <-ch; ok == true {
			fmt.Println("读到数据：", count)
		} else {
			break
		}
	}
}

func GetChannel() {
	// 定义有缓冲通道
	T_ch := make(chan int, 10)
	go func() {
		for i := 0; i <= 10; i++ {
			T_ch <- i
			time.Sleep(time.Second)
		}
	}()
	// 遍历通道
	for data := range T_ch {
		fmt.Println("遍历通道是 -> ", data)
		if data == 3 {
			print("通道值为: 6 退出\n")
			break
		}
	}
}

func Chan_Demo_0() {
	TestChannel()
	GetChannel()
	TestChannel_Status()
}
