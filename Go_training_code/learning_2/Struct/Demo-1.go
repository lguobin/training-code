package Struct

import (
	"fmt"
)

type Books struct {
	title string
	// book_id int
}

func (b Books) AAA() string {
	temp := "方法 run..." + b.title
	return temp
}

func (b *Books) BBB() string {
	temp := "方法 run..." + b.title
	return temp
}

func Struct_Demo_1() {
	fmt.Println("\t --- 结构体实例化 --- ")

	a := Books{"\t测试书名"}
	b := &Books{"BBBB"}
	fmt.Printf("%s -- \n", a.AAA())
	fmt.Printf("%s -- \n", b.BBB())
}
