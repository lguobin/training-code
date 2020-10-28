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
}
