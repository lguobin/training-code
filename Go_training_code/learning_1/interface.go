// interface
package main

import (
	"fmt"
)

// 接口声明、实现多态
type Sleeper interface {
	Sleep()
}
type Eater interface {
	// 声明一个 Eat 方法
	Eat(foodName string)
}
type CCCC interface {
	Sleeper
	Eater
}

type A struct {
	Name string
}
type B struct {
	Name string
}

func (a A) Sleep() {
	fmt.Printf("打印 A函数： %v\n", a.Name)
}
func (b B) Sleep() {
	fmt.Printf("打印 B函数： %v\n", b.Name)
}
func (a A) Eat(foodName string) {
	fmt.Printf("AAA %s is eating %s\n", a.Name, foodName)
}
func (b B) Eat(foodName string) {
	fmt.Printf("AAA %s is eating %s\n", b.Name, foodName)
}
func Temp_Sleep(s Sleeper) {
	// 注意参数是一个 interface
	s.Sleep()
}

// 空接口实现泛型
func _Temp_Print(i interface{}) {
	switch o := i.(type) {
	case int:
		fmt.Printf("整型 -> %d\n", o)
	case float64:
		fmt.Printf("浮点型 -> %f\n", o)
	case string:
		fmt.Printf("字符串 -> %s\n", o)
	case bool:
		fmt.Printf("布尔型 -> %v\n", o)
	case map[string]string:
		fmt.Printf("Map型 -> %v\n", o)

	default:
		fmt.Printf("缺省类型 -> %+v\n", o)
	}
}

func Run_interface_Demo_1() {
	fmt.Println("\t --- Run_interface_Demo_1 ---")
	var _temp Sleeper
	a := A{Name: "xiaobai"}
	b := B{Name: "hellokitty"}
	_temp = a
	Temp_Sleep(_temp) // A 的 Sleep()
	_temp = b
	Temp_Sleep(_temp) // B 的 Sleep()

	_temp_slice := []Sleeper{A{Name: "AAAAAAAA"}, B{Name: "BBBBBBB"}}
	for _, s := range _temp_slice {
		s.Sleep()
	}
}

func Run_interface_Demo_2() {
	fmt.Println("\t --- Run_interface_Demo_2 ---")
	_temp := []CCCC{A{Name: "aaaaaaaaa"}, B{Name: "bbbbbbbbbbbb"}}
	foodName := "零食 - 薯片"
	for _, value := range _temp {
		value.Sleep()
		value.Eat(foodName)
	}
}

func Run_interface_Demo_3() {
	fmt.Println("\t --- Run_interface_Demo_3 ---")
	_temp := []CCCC{A{Name: "A - A"}, B{Name: "B - B"}}
	foodName := "饼干"
	for _, value := range _temp {
		value.Sleep()
		value.Eat(foodName)

		// 判断断言 type assert
		if a, ok := value.(A); ok {
			fmt.Printf("断言: %v 通过\n", a.Name)
		}
		if b, ok := value.(B); ok {
			fmt.Printf("断言: %v 通过\n", b.Name)
		}
	}
}

// 空接口
func Run_interface_Demo_4() {
	fmt.Println("\t --- Run_interface_Demo_4 | 空接口实现 ---")
	_temp := []interface{}{A{Name: "空接口实现泛型 - A"}, B{Name: "空接口实现泛型 - B"}}
	for _, value := range _temp {
		if a, ok := value.(A); ok {
			fmt.Printf("断言: %v 通过\n", a.Name)
		}
		if b, ok := value.(B); ok {
			fmt.Printf("断言: %v 通过\n", b.Name)
		}
	}
}

// 空接口实现泛型
func Run_interface_Demo_5() {
	fmt.Println("\t --- Run_interface_Demo_5 | 空接口实现泛型 ---")
	_Temp_Print(1)
	_Temp_Print(1.2)
	_Temp_Print("123")
	_Temp_Print(true)
	_Temp_Print(map[string]string{"Name": "张三", "Pwd": "123"})
}

// 嵌入接口练习
type USB interface {
	Name() string
	Connect
}
type Connect interface {
	Connect()
}
type PhoneConnect struct {
	name string
}

func (temp PhoneConnect) Name() string {
	return temp.name
}
func (temp PhoneConnect) Connect() {
	fmt.Println("Connect: ", temp.name)
}
func DisConnect(_usb USB) {
	// 判断传入类型
	if pc, ok := _usb.(PhoneConnect); ok {
		fmt.Println("\tif 判断传入类型正常")
		fmt.Println("\t\tDisconnected.", pc.name)
	} else {
		fmt.Println("Unknown Device.")
	}

	// 另一种判断类型方法
	switch v := _usb.(type) {
	case PhoneConnect:
		fmt.Println("\tswitch 判断传入类型正常", v.name)
	default:
		fmt.Println("Unknown Device.")
	}
}
func Run_interface_Demo_6() {
	fmt.Println("\t --- Run_interface_Demo_6 | 嵌入接口 - 练习代码 ---")

	// 第一种方式
	var a USB
	a = PhoneConnect{"第 1 种方式连接 -> 127.0.0.1"}
	a.Connect()
	DisConnect(a)

	// 第二种方式
	b := PhoneConnect{"第 2 种方式连接 -> 0.0.0.1"}
	b.Connect()
	DisConnect(b)

	// 第 3 种方式
	c := PhoneConnect{"第 3 种方式连接 -> 0.0.0.1"}
	var _C Connect
	_C = Connect(c)
	_C.Connect()
	// DisConnect(_C)

}
