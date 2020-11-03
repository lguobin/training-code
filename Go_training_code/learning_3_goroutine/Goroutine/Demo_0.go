package Goroutine

import (
	"fmt"
	"runtime"
)

func Demo_0() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	// 可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
