package main

import (
	"fmt"
)

func Basis_Demo() {
	boolDemo()
	int_basis_Demo()
	testConvert()
	constantDemo()
	arrayDemo()

	sliceDemo()
	test_Slice_Dem()
	mapDemo()
	if_for_switch_Demo()
}

// 所有练习定义函数方法
func Function_Demos() {
	// Go 函数定义 - 开始
	func_sun_Demo(1, []int{1, 2, 3}...)
	func_str_Demo("", 0)

	str_ptr := "『值传递』 -> lao wang"
	str_Demo(str_ptr)
	fmt.Printf("通过参数传递无法修改 -> %s\n", str_ptr)

	ptr_changeString_Demo(&str_ptr)
	fmt.Printf("通过指针传递底层地址修改传参 -> %s\n", str_ptr)

	m := map[string]string{"name": "lao wang"}
	str_Map_Demo(m)
	fmt.Println(m)

	// 调用 - 匿名函数
	test_anonymous_function()
	testFuncType_Demo()
	testMapTpye_Demo()

	testFunction_Demo()

	// 判断奇数
	test_OddNumber_Demo()

	// 闭包
	test_ClosePackage_Demo()

	// 递归函数
	test_Recursion_Demo()
}

func Test_Error_Demo() {
	test_Defer_Demo()

	// 定义null, 并判断返回接口
	if x, err := testError(10, 0); err != nil {
		fmt.Println("被除参数不能传 -> 零")
	} else {
		fmt.Println(x, err)
	}
}

func Struct_Training() {
	fmt.Println("结构体")
	Running_Demo()
	niming_Structure_Demo()
	Struct_method()
}

func Interface_training() {
	fmt.Println("Go 接口")
	Run_interface_Demo_1()
	Run_interface_Demo_2()
	Run_interface_Demo_3()
	Run_interface_Demo_4()
	Run_interface_Demo_5()
	Run_interface_Demo_6()
}

func training() {
	// 程序练习
	key_and_value_change_A()
	defer_Demo()
	Panic_Demo()
	defer_Exercise()
	niming_Structure_Demo()
	// 程序练习 - 结束

	// 可变参数
	fmt.Printf("可变参数传 list 相加之和: %v\n", sum_training([]int{1, 2, 3}...))

	// defer - 主动引发错误练习
	// defer_training()
}

func main() {
	// Basis_Demo()

	// Function_Demos()
	// Test_Error_Demo()

	// Enum_Demo()
	// Test_Counter_Demo()

	// // 序列化
	// Run_Animal()

	// // 结构体
	// Struct_Training()
	Interface_training()

	// training()
}
