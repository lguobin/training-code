package Channel

import (
	"fmt"
	"time"
)

func ch1(ch chan string) {
	for {
		time.Sleep(time.Second * 2)
		ch <- "管道 -  AAA  -"
	}
}

func ch2(ch chan string) {
	for {
		time.Sleep(time.Second * 1)
		ch <- "管道 -  BBB  -"
	}
}

func ch3(ch chan string) {
	for {
		time.Sleep(time.Second * 1)
		ch <- "管道 -  CCC  -"
	}
}

func ch4(ch chan string) {
	for {
		time.Sleep(time.Second * 4)
		ch <- "管道 -  DDDDDDDDDDD  -"
	}
}

func runSelect() {
	A := make(chan string)
	B := make(chan string, 2)
	C := make(chan string, 4)
	D := make(chan string, 8)

	go ch1(A)
	go ch2(B)
	go ch3(C)
	go ch4(D)

	for {
		select {
		case a := <-A:
			fmt.Println("---->   ", a)
		case b := <-B:
			fmt.Println("---->   ", b)
		case c := <-C:
			fmt.Println("---->   ", c)
		case d := <-D:
			fmt.Println("---->   ", d)
		}
	}
}

func timeOut() {
	ch := make(chan int)
	out := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 2)
		out <- true
	}()
	select {
	case <-ch:
	case <-out:
		fmt.Println("没有从 <-ch 中取到数据，此时能从timeout中取得数据")
	}
}

func Chan_Demo_7() {
	// runSelect()
	timeOut()
}
