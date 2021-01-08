package main

import "fmt"

// 可变数量的参数
// more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	fmt.Println(a)
	return a
}

func TestPrint(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	// 匿名函数
	var Add = func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(Add)

	// 可变数量的参数
	more := []int{1, 2, 3, 4}
	Sum(1, more...)

	// 当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果
	var a = []interface{}{123, "abc"}
	TestPrint(a...)
	TestPrint(a)
}
