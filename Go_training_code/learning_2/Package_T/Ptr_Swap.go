package Package_T

import (
	"fmt"
)

func swap(n1 *int, n2 *int) (int, int) {
	temp := *n1
	*n1 = *n2
	*n2 = temp
	fmt.Println("通过指针交换两个数的值", *n1, *n2)

	*n1, *n2 = *n2, *n1
	fmt.Println("再换回来", *n1, *n2)
	return *n1, *n2
}

func Nine_Nine(sum int) {
	fmt.Println("打印 - 九九乘法表")
	for i := 1; i <= sum; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func printPyramid(Level int) {
	fmt.Println("打印 - 空心金字塔")
	for i := 1; i <= Level; i++ {
		for k := 1; k <= Level-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == Level {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Ptr_Swap() {
	a, b := 1, 2
	_, c := swap(&a, &b)
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)

	Nine_Nine(9)
	printPyramid(20)
}
