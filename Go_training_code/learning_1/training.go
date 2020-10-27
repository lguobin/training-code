package main

import (
	"fmt"
)

// 字典key和value对调
func key_and_value_change_A() {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Printf("字典key和value对调 | %v\n", m2)
}

// defer 匿名函数
func defer_Demo() {
	for i := 0; i < 5; i++ {
		// defer fmt.Println(i)
		defer func() {
			fmt.Println(i)
		}()
	}
}

// 程序恢复执行: Panic | recover
func Panic_Demo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("恢复程序...")
		}
	}()
	panic("吓到程序停止")
}

// Exercise defer - 练习
func defer_Exercise() {
	fmt.Println("\tExercise defer - 练习")
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)
		defer func() { fmt.Println("defer_closure i =", i) }()
		fs[i] = func() { fmt.Println("closure i =", i) }
	}
	for _, f := range fs {
		f()
	}
	fmt.Println("\tExercise defer - 练习 - End")
}

// 可变参数
func sum_training(num_list ...int) int {
	temp := 0
	for _, valus := range num_list {
		temp += valus
	}
	return temp
}

// defer - start
// func LogOut() {
// 	fmt.Println("LogOut...")
// }
// func defer_training() {
// 	defer LogOut()
// 	fmt.Println("登录用户...")
// 	panic("主动引发错误")
// 	fmt.Println("不会输出")
// }
// defer - End
