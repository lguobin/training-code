package Package_T

import (
	"fmt"
	"strconv"
	"strings"
)

func String_Demo() {
	fmt.Println("\n\t --- 类型转换 --- ")
	str := "hello北"
	fmt.Println("str len=", len(str)) // 8

	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c\n", r[i])
	}

	// 字符串转整数:	 n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("65")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("字符串转整数 -> ", n)
	}

	// 4) 整数转字符串  str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str= %v, str= %T\n", str, str)

	// 5) 字符串 转 []byte:  var bytes = []byte("ABCDE")
	var bytes = []byte("ABCDE")
	fmt.Printf("bytes = %v\n", bytes) // ASCII 65 ~ 69

	// 6) []byte 转 字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{65, 66, 67})
	fmt.Printf("str = %v\n", str) // ABC

	// 10进制转 2, 8, 16进制:  str = strconv.FormatInt(123, 2),返回对应的字符串
	// 只需要修改 strconv.FormatInt(整型参数, 转进制参数)
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的 2 进制是 = %v\n", str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("123对应的 8 进制是 = %v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的 16 进制是 = %v\n", str)

	//统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "e") // 4
	num := strings.Count("ceheese", "e")
	fmt.Printf("e 出现过: %v次\n", num)
}
