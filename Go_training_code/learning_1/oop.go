// oop

// 封装 \ 继承 \ 多态

package main

import (
	"fmt"
)

type Enum int

// const (
// 	// 这里如果是自增用 iota 更好
// 	Init    Enum = 0
// 	Success Enum = 1
// 	Fail    Enum = 2
// )

const (
	Init Enum = iota
	Pass
	Fail

	// 枚举对应的中文
	InitName    = "初始化"
	SuccessName = "成功"
	FailName    = "失败"
)

func (e Enum) Int() int {
	//返回时候强转 int 类型，可以直接用
	return int(e)
}

func (e Enum) _String() string {
	return []string{
		InitName,
		SuccessName,
		FailName,
	}[e]
}

func Enum_Demo() {
	status := 0

	// fmt.Println(Init == status)
	// ↑↑↑ (mismatched types Enum and int)
	fmt.Printf("调用 Init 需要转换, 当前类型: %T | 值: %v\n", Init, int(Init) == status)
	fmt.Printf("调用 Init.Int(), 当前类型: %T | 值: %v\n", Init.Int(), Init.Int() == status)

	status_name := Pass // 调用自己定义的类型
	fmt.Println(status_name._String())
}

// 定义一个 Counter 类型
func Test_Counter_Demo() {
	type Counter map[string]int
	c := Counter{}
	c["Work"]++
	fmt.Println(c)

	type Queue []int
	q := make(Queue, 0)
	q = append(q, 1)
	fmt.Println(q)
}

// 源码延伸
/*
通过看一些 go 的源码，我们可以学习并且模仿 go 的惯用法，
比如本文提到的 Enum 类型，
在 go 的源码你可以找到类似实现。
以下是 go 的内置的 http server 中关于枚举的实现方式(去掉了注释)：
*/

/*
type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

var stateName = map[ConnState]string{
	StateNew:      "New",
	StateActive:   "Active",
	StateIdle:     "Idle",
	StateHijacked: "Hijacked",
	StateClosed:   "Closed",
}

func (c ConnState) Test_String() string(
	return stateName[c]
)

*/
