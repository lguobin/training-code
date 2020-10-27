// function
package main

import "fmt"

func func_sun_Demo(a int, b ...int) (res int) {
	for _, val := range b {
		a += val
	}

	fmt.Println("结果打印 -> ", res)
	return
}

func func_str_Demo(a string, b int) string {
	if a == "" {
		a = "default-a"
	}
	if b == 0 {
		b = 1
	}
	c := fmt.Sprintf("%s - %d", a, b)
	fmt.Println(c)
	return c
}

func str_Demo(s string) {
	s = "值传递"
	fmt.Println(s)
}

func str_Map_Demo(m map[string]string) {
	m["job"] = "老司机"
}

func ptr_changeString_Demo(s *string) {
	*s = "New Message"
}

// 匿名函数
func test_anonymous_function() {
	fmt.Println("定义匿名函数 - 开始")
	func(s string) {
		fmt.Println(s)
	}("匿名函数 - 结束")
}

func testFuncType_Demo() {
	myPrint := func(s string) { fmt.Println(s) }
	myPrint("测试 - 函数类型")
}

func testMapTpye_Demo() {
	funcMap := map[string]func(int, int) int{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
	}
	fmt.Printf("相加: %d\n", funcMap["add"](3, 2))
	fmt.Printf("相减: %d\n", funcMap["sub"](5, 4))
}

// 把函数当成参数传递到另一个函数里面
func Double(x int) int {
	return x * 2
}

func Apply(x int, f func(int) int) int {
	return f(x)
}
func testFunction_Demo() {
	fmt.Printf("把函数当成参数传递进去 -> %d\n", Apply(2, Double))
	fmt.Printf("把函数当成参数传递进去 -> %d\n", Apply(4, Double))
}

// 高阶函数 - 判断所有奇数(odd number)
func FilterIntSlice(intVals []int, param func(i int) bool) []int {
	res := make([]int, 0)
	for _, val := range intVals {
		if param(val) {
			res = append(res, val)
		}
	}
	return res
}

func test_OddNumber_Demo() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	isOdd := func(i int) bool { return i%2 != 0 }
	fmt.Printf("打印所有奇数 -> %d\n", FilterIntSlice(ints, isOdd))
}

// 闭包
func test_ClosePackage_Demo() {
	str_ := ".go"
	addStr := func(name string) string {
		return name + str_
	}
	fmt.Printf("打印闭包 -> %s\n", addStr("测试文件"))
	fmt.Println("到 goroutine 的时候，我们会看到一个 for 循环里使用闭包的坑。")
}

// 递归函数
func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}

func test_Recursion_Demo() {
	fmt.Println("《 测试递归函数 》")
	fmt.Println("在线数学工具 -> https://zh.numberempire.com/")
	fmt.Println(Fib(10))
}
