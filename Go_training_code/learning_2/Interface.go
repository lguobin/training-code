// Interface
package main

import (
	"fmt"
)

func Interface__init__() {
	fmt.Println("\t接口 interface")
}

// 运输方式 - 类型
type Transporter interface {
	BicycleTran()
	CarTran()
}

// 驾驶员 - 结构
type Driver struct {
	Name        string
	Driver_type string
}

// 实现运输方式 - 方法
func (temp *Driver) BicycleTran() {
	fmt.Println("共享单车运输...")
}
func (temp *Driver) CarTran() {
	fmt.Println("共享汽车运输...")
}

// 只要实现了 Transporter接口的类型都可以作为参数
func trans(t Transporter) {
	t.BicycleTran()
}

// 多个类型共用一个接口
type Service interface {
	Start()
	Log(string)
}
type Logger struct {
}
type GameService struct {
	Logger
}

func (log *Logger) Log(s string) {
	fmt.Println("日志输出:", s)
}
func (gome *GameService) Start() {
	fmt.Println("游戏服务启动")
}

// 多个类型共用一个接口 - End
//
// - End

func Bsisc_Interface() {
	Interface__init__()

	task := &Driver{"单车老四", "两厢通风自行车"}
	trans(task)

	run_game := new(GameService)
	run_game.Start()
	run_game.Log("正在打开游戏...")

	m := make(map[string]interface{})
	m["Name"] = "空接口名"
	m["Page"] = 1
	m["total"] = 1
	fmt.Println("利用空接口，可以实现任意类型的存储: ", m)
}
