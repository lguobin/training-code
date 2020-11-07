package Goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

var ticket = 4

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
	}
}

func Demo_6() {
	fmt.Println("\n\t --- 不带锁多协程处理")
	go saleTickets("售票口 - A ")
	go saleTickets("售票口 - B ")
	go saleTickets("售票口 - C ")
	go saleTickets("售票口 - D ")
	time.Sleep(time.Second * 3)
}
