package Goroutine

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg_9 *sync.WaitGroup

func Write_Data(i int) {
	defer wg_9.Done()
	fmt.Println(i, "开始写：write start。。")
	// 写操作上锁
	rwMutex.Lock()
	fmt.Println(i, "正在写：writing。。。。")
	time.Sleep(time.Second * 3)
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over。。")
}

func Demo_9() {
	fmt.Println("\n\n\t --- 互斥锁 -> RWMutex(读写锁)")
	rwMutex = new(sync.RWMutex)
	wg_9 = new(sync.WaitGroup)

	wg_9.Add(3)
	go Write_Data(1)
	go func(i int) {
		defer wg_9.Done()
		fmt.Println(i, "开始读：read start。。")

		// 读操作上锁
		rwMutex.RLock()
		fmt.Println(i, "正在读取数据：reading。。。")
		time.Sleep(time.Second * 3)
		// 读操作解锁
		rwMutex.RUnlock()
		fmt.Println(i, "读结束：read over。。。")
	}(2)
	go Write_Data(3)

	wg_9.Wait()
	fmt.Println("\t --- main，解除阻塞。。")
}
