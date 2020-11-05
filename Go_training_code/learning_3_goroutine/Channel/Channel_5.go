package Channel

import (
	"fmt"
	"time"
)

func timeMore(ch chan string) {
	ch <- "任务"
	fmt.Println("模拟耗时操作: --- 限制并发 ---")
	time.Sleep(time.Second)
	<-ch
}

func Chan_Demo_5() {
	ch := make(chan string, 5)
	fmt.Println("\t --- 现在有100个并发，限制为5个并行 --- ")
	for i := 0; i < 20; i++ {
		go timeMore(ch)
	}
	for {
		time.Sleep(time.Microsecond)
	}
}
