package Channel

import (
	"fmt"
)

type all_T struct {
	Name string
	Age  int
}

func Interface_chan() {
	// 可以存放任意数据类型管道
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "测试"
	allChan <- all_T{"测试数据", 55}
	<-allChan
	<-allChan
	get_data := <-allChan
	fmt.Printf("get_data -> %T\n", get_data)

	change_struct := get_data.(all_T)
	fmt.Printf("把管道接口数据转换回结构体数据 -> %s - %d\n", change_struct.Name, change_struct.Age)
}
