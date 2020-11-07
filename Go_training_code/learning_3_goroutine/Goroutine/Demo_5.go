package Goroutine

import (
	"fmt"
	"time"
)

func Demo_5() {
	fmt.Println("协程会共享内存，[ a ]会改变，time.Sleep上移会输出: a = 3")
	a := 1
	go func() {
		a = 2
		fmt.Println("子goroutine。。", a)
	}()
	a = 3
	time.Sleep(time.Millisecond * 500)
	fmt.Println("main goroutine。。", a)
}
