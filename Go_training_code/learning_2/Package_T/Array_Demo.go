package Package_T

import (
	"fmt"
	"math/rand"
	"time"
)

func New_array() {
	fmt.Println("\t\t四种初始化数组的方式")
	// 四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Printf("numArr01的地址 = %p\n\tnumArr01[0] 地址 %p\n\tnumArr01[1] 地址 %p\n\tnumArr01[2] 地址 %p\n",
		&numArr01, &numArr01[0], &numArr01[1], &numArr01[2])

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)
	//这里的 [...] 是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=", numArr04)

	//类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05=", strArr05)
}

func array_function(temp []int) {
	fmt.Println("打印数组参数 -> ", temp)
}
func array_function_2(temp ...int) {
	fmt.Println("打印不固定参数 -> ", temp)
}
func array_function_3(temp *[]int) {
	fmt.Printf("temp 指针的地址 = %p | %v\n", &temp, temp)
	(*temp)[0] = 666666
	fmt.Printf("temp 指针的地址 = %p | %v\n", &temp, temp)
}

func array_abc() {
	var chr [26]byte
	for i := 0; i < 26; i++ {
		chr[i] = 'A' + byte(i)

	}
	for i := 0; i < 26; i++ {
		fmt.Printf("打印 ASCII -> %c \n", chr[i])
	}
}

func arryay_Get_max() {
	Arr := [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := Arr[0]
	max_num := 0
	for i := 0; i < len(Arr); i++ {
		if maxVal < Arr[i] {
			maxVal = Arr[i]
			max_num = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n\n", maxVal, max_num)
}

func array_random() {
	var Arr [5]int
	temp := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(Arr); i++ {
		Arr[i] = rand.Intn(100)
	}

	fmt.Println("交换前~=", Arr)
	for i := 0; i < len(Arr)/2; i++ {
		temp = Arr[len(Arr)-1-i]
		Arr[len(Arr)-1-i] = Arr[i]
		Arr[i] = temp
	}
	fmt.Println("交换后~=", Arr)
}

func Slice_fib(n int) []uint16 {
	fmt.Println("切片 - 斐波那契数列")
	fibSlice := make([]uint16, n)
	fibSlice[0], fibSlice[1] = 1, 1
	for i := 2; i < n; i++ {
		fibSlice[i] = fibSlice[i-1] + fibSlice[i-2]
	}
	return fibSlice
}
func Slice_Demo() {
	var intArr [5]int = [...]int{1, 22, 33, 44, 55}
	slice := intArr[0:3]
	fmt.Println("slice 的元素是 =", slice)       //  1, 22, 33
	fmt.Println("slice 的元素个数 =", len(slice)) // 3
	fmt.Println("slice 的容量 =", cap(slice))   // 切片的容量是可以动态变化

	fmt.Printf("intArr[0]的地址=%p | 值: %v\n", &intArr[0], intArr[0])
	fmt.Printf("slice[0]的地址=%p  | 值: %v\n", &slice[0], slice[0])

	// 修改切片，原数组也会变更值
	slice[1] = 6666666
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)

	// 用append内置函数，可以对切片进行动态追加
	slice2 := append(slice, 400, 500, 600)
	fmt.Println("slice 的元素是 =", slice)
	fmt.Println("slice2 的元素是 =", slice2)

	// 用copy内置函数完成拷贝
	slice3 := make([]int, 10)
	copy(slice3, slice2)
	fmt.Println("slice2 的元素是 =", slice2, "\nslice3 的元素是 =", slice3)
}

func Slice_to_String() {
	fmt.Println("\t\tstring底层是一个byte数组，因此string也可以进行切片处理")
	str := "abc@atguigu"
	slice := str[6:]
	fmt.Println(slice)

	//string是不可变的，也就说不能通过 str[0] = 'z' 方式来修改字符串
	//str[0] = 'z' [编译不会通过，报错，原因是string是不可变]
	//如果需要修改字符串，可以先将string -> []byte / 或者 []rune -> 修改 -> 重写转成string
	// "hello@atguigu" =>改成 "zello@atguigu"

	temp := []byte(str)
	fmt.Printf("修改前: %v\n", string(temp))
	temp[0], temp[1], temp[2] = 'A', 'B', 'C'
	fmt.Printf("修改后: %v\n", string(temp))

	arr1 := []rune(str)
	arr1[0] = '北'
	str = string(arr1)
	fmt.Println("str=", str)
}

func Array_Demo() {
	fmt.Println("\t --- 数组 --- ")
	New_array()

	a := []int{1, 2, 3}
	array_function(a)
	array_function([]int{1, 2, 3})
	array_function_2([]int{1, 2, 3}...)
	array_function_2(a...)
	array_function_3(&a)

	array_abc()
	var bb [3]byte
	var cc [3]rune
	bb[0] = 'A'
	cc[0] = 'A'
	fmt.Println(bb, cc)

	arryay_Get_max()
	array_random()

	fmt.Println(Slice_fib(10))
	Slice_Demo()
	Slice_to_String()

}
