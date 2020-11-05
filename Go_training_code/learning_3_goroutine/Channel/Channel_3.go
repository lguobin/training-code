package Channel

import (
	"fmt"
	"time"
)

func Chan_Demo_3() {
	intChan := make(chan int, 10)
	stringChan := make(chan string, 5) // 定义一个管道 5个数据string
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	for i := 0; i < 5; i++ {
		stringChan <- "hello - " + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 问题，在实际开发中，可能我们不好确定什么关闭该管道.
	// 可以使用select 方式可以解决
	// label:
	for {
		select {
		//注意: 这里，如果intChan一直没有关闭，不会一直阻塞而deadlock, 会自动到下一个case匹配
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据 ->  %d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据 ->  %s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到了，不玩了, 程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			return
		}
	}
	//break label
}
