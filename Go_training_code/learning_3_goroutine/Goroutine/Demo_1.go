package Goroutine

import (
	"fmt"
	"strconv"
	"time"
)

func tesst() {
	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () 打印输出 -> " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func Demo_1() {
	// 开启一个协程
	go tesst()

	// 主进程运行
	for i := 1; i <= 12; i++ {
		fmt.Println("main() 打印输出 -> " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
