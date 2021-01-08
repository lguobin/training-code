package main

import "fmt"

// 定义函数类型
type funcCaller func(interface{})

// 实现 Invoker 的 Call
func (this funcCaller) Call(task interface{}) {
	// 调用函数
	this(task)
}

// 调用接口
type invoker interface {
	Call(interface{})
}

func InterFace_B() {
	// 声明接口类型
	var inv invoker
	// 将匿名函数转为 funcCaller 类型
	inv = funcCaller(func(v interface{}) {
		fmt.Println("打印内容: ", v)
	})
	// 使用接口调用 funcCaller.Call, 内部会调用函数本体
	inv.Call(" - 测试信息 → 内部会调用函数本体")
}
