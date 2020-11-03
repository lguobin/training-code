package Goroutine

import (
	"fmt"
	"time"
)

// 向 intChan 存放入 1-8000个数
func putMath(intChan chan int) {
	for i := 0; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就
// 放入到 primeChan
func primeMath(intChan chan int, primeChan chan int, exitChan chan bool) {
	var result bool
	for {
		count, ok := <-intChan
		if !ok {
			fmt.Println("intChan 取不到..")
			break
		}
		result = true
		for i := 2; i < count; i++ {
			// 说明该count不是素数
			if count%i == 0 {
				result = false
				break
			}
		}
		if result {
			primeChan <- count
		}
	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	// 这里我们还不能关闭 primeChan
	// 向 exitChan 写入true
	exitChan <- true
}
func Demo_3() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) // 放入结果
	exitChan := make(chan bool, 8)     // 标识退出的管道 - 4个
	Coroutine := 8                     // 协程数量

	start := time.Now().Unix()
	go putMath(intChan)
	// 开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就
	// 放入到primeChan
	for i := 0; i < Coroutine; i++ {
		go primeMath(intChan, primeChan, exitChan)
	}

	// 主线程
	go func() {
		for i := 0; i < Coroutine; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时 = ", end-start)
		close(primeChan)
	}()

	temp := []int{}
	for {
		count_end, ok := <-primeChan
		temp = append(temp, count_end)
		if !ok {
			break
		}
	}
	// fmt.Printf("素数列表 = %d\n", temp)
	fmt.Println("main线程退出")
}
