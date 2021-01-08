package main

import "fmt"

// 定义结构体
type structCaller struct{}

func (this *structCaller) Call(p interface{}) {
	fmt.Println("打印信息: ", p)
}

type invoker_c interface {
	Call(interface{})
}

func InterFace_C() {
	// 声明接口变量
	var inv invoker_c
	var str structCaller

	// 实例化结构体传结构内存地址到接口
	inv = &str
	inv.Call("通过传递结构内存地址实例化输出")
}
