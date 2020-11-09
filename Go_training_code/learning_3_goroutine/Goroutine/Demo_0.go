package Goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func T_Gosched() {
	text := "这个函数的作用是让当前 `goroutine` 让出 `CPU`，当一个 `goroutine` 发生阻塞，`Go` 会自动地把与该 `goroutine` 处于同一系统线程的其他 `goroutine` 转移到另一个系统线程上去，以使这些 `goroutine` 不阻塞。\n"
	fmt.Println("\n\t ---> 让当前线程让出 `cpu` 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行\n", text)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine。。。")
		}
	}()
	for i := 0; i < 4; i++ {
		runtime.Gosched()
		fmt.Println("主进程 ...")
	}
}
func fun_Test() {
	defer fmt.Println("defer。。。")
	// 终止所在的协程
	runtime.Goexit()
	fmt.Println("fun_Test 函数。。。\n不会被打印出来")
}
func T_Goroutine() {
	fmt.Println("\n\t ---> 终止协程\\退出当前 `Goroutine`(但是`defer`语句会照常执行)")

	go func() {
		fmt.Println("goroutine开始。。。")
		fun_Test()
	}()
	time.Sleep(time.Millisecond * 1000)
}

func T_Goroutine_B() {
	fmt.Println("\n\t ---> 终止协程\\退出当前 `Goroutine`(但是`defer`语句会照常执行)")

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

func Demo_0() {

	fmt.Println("G获取goroot目录-->", runtime.GOROOT())
	fmt.Println("获取操作系统-->", runtime.GOOS)

	cpuNum := runtime.NumCPU() //获取当前系统的CPU核心数
	runtime.GOMAXPROCS(cpuNum) //Go中可以轻松控制使用核心数
	fmt.Println("cpuNum=", cpuNum)

	// 可以自己设置使用多个cpu
	fmt.Println("ok")

	T_Gosched()
	T_Goroutine()
	T_Goroutine_B()
}
