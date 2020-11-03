package Goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成

/*
// 阶乘: 结果 = 1*2*3...*n
// Python:
	for i in range(1, n + 1):
		factorial = factorial*i
		或:
		factorial *= i
*/

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做出一个全局的.

var (
	MAP = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁，
	// sync 是包: synchornized 同步
	// Mutex : 是互斥
	lock sync.Mutex
)

func factorial(x int) {
	res := 1
	for i := 1; i <= x; i++ {
		// res = res * i
		res *= i
	}

	lock.Lock()
	MAP[x] = res // 并发映射写
	lock.Unlock()
}

func Demo_2() {
	// 我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 10; i++ {
		go factorial(i)
	}

	time.Sleep(time.Second * 2)

	lock.Lock()
	for i, v := range MAP {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
