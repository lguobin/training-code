package Goroutine

import (
	"fmt"
	"time"
)

func running() {
	var count int
	for {
		count++
		fmt.Println(" >>", count)
		time.Sleep(time.Millisecond * 1000)
	}
}

func running_1() {
	for i := 0; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
}

func Demo_4() {
	// go running()
	// var input string
	// fmt.Scanln(&input)
	running_1()
}
