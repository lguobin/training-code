package Goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func Demo_0() {

	cpuNum := runtime.NumCPU() //获取当前系统的CPU核心数
	runtime.GOMAXPROCS(cpuNum) //Go中可以轻松控制使用核心数
	fmt.Println("cpuNum=", cpuNum)

	// 可以自己设置使用多个cpu
	fmt.Println("ok")

	for i := 0; i <= 5; i++ {
		defer fmt.Println("defer ", i)
		go func(i int) {
			if i == 1 {
				runtime.Goexit()
			}
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}
