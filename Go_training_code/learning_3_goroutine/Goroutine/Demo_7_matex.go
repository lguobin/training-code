package Goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket_lock = 10
var wg_7 sync.WaitGroup
var matex sync.Mutex // 创建锁头

func saleTickets_Lock(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg_7.Done()
	for {
		matex.Lock()
		if ticket_lock > 0 {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			fmt.Println(name, "售出：", ticket_lock)
			ticket_lock--
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			matex.Unlock()
			break
		}
		matex.Unlock()
	}
}

func Demo_7() {
	fmt.Println("\n\t --- 带锁多协程处理")
	wg_7.Add(4)
	go saleTickets_Lock("售票口 - A ")
	go saleTickets_Lock("售票口 - B ")
	go saleTickets_Lock("售票口 - C ")
	go saleTickets_Lock("售票口 - D ")
	wg_7.Wait()
}
