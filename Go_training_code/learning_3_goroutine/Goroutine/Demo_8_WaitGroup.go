package Goroutine

import (
	"fmt"
	"sync"
)

var wg_8 sync.WaitGroup

func Demo_8() {
	wg_8.Add(2)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("fun1.。。i:", i)
		}
		// 给wg等待中的执行的goroutine数量减1.同Add(-1)
		wg_8.Done()
	}()

	go func() {
		for j := 0; j <= 10; j++ {
			fmt.Println("\tfun2..j,", j)
		}
		// 给wg等待中的执行的goroutine数量减1.同Add(-1)
		defer wg_8.Done()
	}()

	fmt.Println("\n\t --- main进入阻塞状态。。。等待wg中的子goroutine结束。。")
	// 表示main goroutine 进入等待，意味着阻塞
	wg_8.Wait()
	fmt.Println("main，解除阻塞。。")
}
