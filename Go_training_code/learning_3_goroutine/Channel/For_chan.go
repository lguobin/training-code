package Channel

import (
	"fmt"
)

// 写数据
func writedata(intchan chan int) {
	for i := 0; i < 50; i++ {
		intchan <- i
	}
	close(intchan)
}

func readdata(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Println("读到的数据 -> ", v)
	}
	exitchan <- true
	close(exitchan)
}

func For_chan() {
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)
	go writeData(intchan)
	go readdata(intchan, exitchan)

	// time.Sleep(time.Millisecond * 500)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}
