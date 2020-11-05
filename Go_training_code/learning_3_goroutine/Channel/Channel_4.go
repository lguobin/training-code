package Channel

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("We Say Hello World!")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var MAP map[int]string
	MAP = make(map[int]string, 10)
	MAP[0] = "TTTTTTTTTTTT"
	fmt.Println(MAP)
}

func Chan_Demo_4() {
	go test()
	go sayHello()
	for i := 0; i < 10; i++ {
		fmt.Println("主函数 ok -> ", i)
		time.Sleep(time.Second)
	}
}
