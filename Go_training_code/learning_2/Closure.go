// Closure
package main

import (
	"fmt"
)

func Closure__init__() {
	fmt.Println("\t闭包")
}
func fn_A(a int) func(temp int) int {
	return func(temp int) int {
		// print(&a, a, "\n")
		return a
	}
}
func sum(value int) func() int {
	return func() int {
		value += 10
		// print(value, "\n")
		return value
	}
}

func Bsisc_Closure() {
	Closure__init__()
	_a, _b := fn_A(1), fn_A(2)

	fmt.Printf("第 1 层运行结果: %d\n", _a(1))
	fmt.Printf("第 1 层运行结果: %d\n", _a(1))

	fmt.Printf("第 2 层运行结果: %d\n", _b(2))
	fmt.Printf("第 2 层运行结果: %d\n", _b(2))

	Add := sum(1)
	fmt.Println(Add())
	fmt.Println(Add())
}
