// basis
package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func int_basis_Demo() {
	fmt.Printf("int32最大值: %v\n", math.MaxInt32)
	// fmt.Printf("int64最大值: %v\n", math.MaxInt64)
	fmt.Printf("Float64最大值: %v\n", math.MaxFloat64)

	var a int8 = 127
	b := 100
	fmt.Printf("%T 最小值: %v & 最大值: %v\n", ^int(^uint(0)>>1), int(^uint(0)>>1))
	fmt.Printf("a 类型: %T, 取值范围: %v\n", a, a)
	fmt.Printf("b 类型: %T, 取值范围: %v|%T\n", b, b, unsafe.Sizeof(b))

	var c int16 = 130
	fmt.Printf("高位转低位,精度丢失: %v\n", int8(c))

	cc := 1129.6
	fmt.Printf("乘法精度丢失: %v\n", cc*100)
	fmt.Println("乘法精度丢失: ", cc*100)

	// 进制转换 - 输出
	var d int8 = 17
	fmt.Printf("十进制输出: %d\n", d)
	fmt.Printf("以二进制输出: %b\n", d)
	fmt.Printf("以八进制输出: %o\n", d)
	fmt.Printf("以十六进制输出: %x\n", d)

	pi := math.Pi
	fmt.Printf("默认输出小数点6位, %f\n", pi)
	fmt.Printf("输出小数点后2位, %.2f\n", pi)

	// uint8 类型相当于是 byte , 代表了ACII码的一个字符
	// ASCII码值
	var A_ascii byte = 'A'
	var B_ascii uint8 = 'B'
	fmt.Printf("输出ASCII码值: %d\n", A_ascii)
	fmt.Printf("输出ASCII码值: %c\n", A_ascii)
	fmt.Printf("输出ASCII码值: %d\n", B_ascii)
	fmt.Printf("输出ASCII码值: %c\n", B_ascii)

	// 字符串拼接
	s1 := `第一行...
	第二行...`
	s2 := "\"Golang !\""
	fmt.Println(s1 + s2)
}

func testConvert() {

	// int -> string
	sint := strconv.Itoa(65)
	fmt.Println(sint)

	// byte -> string
	bytea := byte(1)
	bint := strconv.Itoa(int(bytea))
	fmt.Println(bint)

	// int64 -> string
	sint64 := strconv.FormatInt(int64(97), 10)
	fmt.Println(sint64, sint64 == "97")

	// int64 -> string (hex) ，十六进制
	sint64hex := strconv.FormatInt(int64(97), 16)
	fmt.Println(sint64hex, sint64hex == "61")

	// string -> int
	_int, _ := strconv.Atoi("65")
	fmt.Println(_int, _int == int(65))

	// string -> int64
	_int64, _ := strconv.ParseInt("65", 10, 64)
	fmt.Println(_int64, _int64 == int64(65))

	// string -> int32
	_int32, _ := strconv.ParseInt("65", 10, 32)
	fmt.Println(_int32, int32(_int32) == int32(65))
	fmt.Printf("_int32 -> ASCII: %c\n", _int32)

	// int32 -> string
	_int32str := 65
	strconv.FormatInt(int64(_int32str), 10)
	strconv.Itoa(int(_int32str))
	fmt.Sprint(_int32str)

	// int -> int64 ，不会丢失精度
	var n int = 65
	fmt.Println(int64(n), int64(n) == int64(65))

	//string -> float32/float64
	f := fmt.Sprintf("%.4f", math.Pi)
	// fmt.Println(f) -> 3.1416
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s) // 3.1416
	}

}

func constantDemo() {
	const (
		SunDay = iota + 1
		MonDay
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	if Saturday == 7 {
		fmt.Printf("通过判断, Saturday: %v\n", Saturday)
	}
}

func boolDemo() {
	var a bool
	b := false
	c := true
	if a {
		fmt.Println("a is True")
	}
	if b {
		fmt.Println("b is False")
	}
	if c {
		fmt.Println("c is True")
	}
	fmt.Printf("a: %v, b: %v, c: %v", a, b, c)
}

func arrayDemo() {
	// 数组的长度是类型的一部分
	var arr1 [3]int
	var arr2 [4]string
	fmt.Printf("%T, %T \n", arr1, arr2)

	// 数组的初始化 第一种方法
	var arr3 [3]int
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3
	fmt.Println(arr3)

	// 第二种初始化数组的方法
	var arr4 = [4]int{10, 20, 30, 40}
	fmt.Println(arr4)

	// 第三种数组初始化方法，自动推断数组长度
	var arr5 = [...]int{1, 2}
	fmt.Println(arr5)

	// 第四种初始化数组的方法，指定下标
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}

	for _, value := range a {
		fmt.Print(value, " ")
	}

	fmt.Println()
	// 值类型 引用类型
	// 基本数据类型和数组都是值类型
	var aa = 10
	bb := aa
	aa = 20
	fmt.Println(aa, bb)

	// 数组
	var array1 = [...]int{1, 2, 3}
	array2 := array1
	array2[0] = 3
	fmt.Println(array1, array2)

	// 切片定义
	var array3 = []int{1, 2, 3}
	array4 := array3
	array4[0] = 3
	fmt.Println(array3, array4)

	// 二维数组
	var array5 = [...][2]int{{1, 2}, {2, 3}}
	for i := 0; i < len(array5); i++ {
		for j := 0; j < len(array5[0]); j++ {
			fmt.Println(array5[i][j])
		}
	}

	for _, item := range array5 {
		for _, item2 := range item {
			fmt.Println(item2)
		}
	}
}

func sliceDemo() {
	names := []string{"zhang", "li", "wang", "zhao"}
	_name_1 := names[0]
	names[0] = "1111"
	_name_1 = "222"
	fmt.Println(names, _name_1)
	// 你发现 names 改变会引起 _name_1 也改变，这里起始它们共用了底层结构，注意这个问题

	vals := make([]int, 0)
	for i := 0; i < 10; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)

	// 必须使用省略号的方式『解包』一个 slice 来连接两个 slice
	vals2 := []int{3, 4, 5, 6}
	newvals := append(vals, vals2...)
	fmt.Println(newvals)
}

func test_Slice_Dem() {
	max_size := make([]int, 1000000)

	// bad way
	a := make([]int, 0)
	for _, val := range max_size {
		// 扩容 a 会导致重新分配内存
		a = append(a, val+val)
	}
	fmt.Println(len(a))

	// good way
	b := make([]int, len(max_size))
	for i, val := range max_size {
		// 注意这里是赋值了，不是 append
		b[i] = val + val
	}
	fmt.Println(len(b))

	/*
		课后问题:
			- 请给一个 slice 反向排序？不知道的话请搜索 go 的 sort 文档
			- 什么情况下我们要去关心 slice 的容量呢？append 之后它的容量如何变化呢？
			- sort 包里的稳定排序和非稳定排序有什么区别？
	*/
}

func mapDemo() {
	m := make(map[string]string)
	m["name"] = "张三"
	m["job"] = "老司机"
	m["age"] = "18"
	m["temp"] = "temp_work"

	// 输出一个存在的 key 和一个不存在的 key
	fmt.Printf("打印存在:%v, 打印不存在 key 不会报错:%v\n", m["name"], m["not_found"])
	delete(m, "temp_not_found")
	fmt.Println(m)

	if value, ok := m["job"]; ok {
		fmt.Printf("判断打印 -> m[\"job\"] is %s\n", value)
	} else {
		fmt.Println("字典 key 不存在")
	}

	// 同样使用 for/range 遍历，NOTE：遍历 map 返回的顺序是随机的，不要依赖 map 遍历的顺序
	for k, v := range m {
		fmt.Printf("遍历打印 -> m[\"%s\"]: %s\n", k, v)
	}

	// 如果只需要 k 或者 v 你可以使用 下划线作为占位符忽略值
	for _, v := range m {
		fmt.Println(v)
	}

	// Map -> Set
	map_set := make(map[string]bool)
	map_set["激活"] = true
	map_set["信号"] = true
	map_set["通讯"] = false
	key := "通讯"
	if _, ok := map_set[key]; ok {
		fmt.Println("找到该 key 打印值 -> : ", key)
	}

	// 循环判断
	for k, v := range map_set {
		if v == false {
			fmt.Println(v)
		}
		fmt.Println(k, v)
	}
}

func if_for_switch_Demo() {
	ok := true
	if ok {
		fmt.Println("这是真的!")
	}

	day := 6
	if day == 6 {
		fmt.Println("明天不上班呀!")
	} else if day == 7 {
		fmt.Println("周末好快")
	} else {
		fmt.Println("干活啦")
	}
	m := make(map[string]string)
	m["name"] = "张三"
	m["job"] = "老司机"
	if v, ok := m["name"]; ok {
		fmt.Println(v)
	}

	switch day {
	case 0, 6:
		fmt.Println("switch 输出 -> 周末")
	case 1, 2, 3, 4, 5:
		fmt.Println("switch 输出 -> 工作日")
	default:
		fmt.Println("switch 输出 -> Null")
	}

	a, b := 1, 2
	a, b = b, a
	switch {
	case a > b:
		fmt.Println("switch 输出 -> a > b")
	case a < b:
		fmt.Println("switch 输出 -> a < b")
	}
	int_slice := []int{1, 2, 3}
	for index, item := range int_slice {
		fmt.Printf("切片输出 -> %v - %v\n", index, item)
	}
	for index := range int_slice {
		fmt.Printf("切片输出 -> %v\n", index)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		if k == "job" {
			fmt.Println("判断 K 如果是 Job 直接退出", k)
			break
		} else {
			fmt.Println("继续循环")
			continue
		}
	}
}
