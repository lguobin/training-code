package Package_T

import (
	"fmt"
)

func Error_Demo() {
	fmt.Println("\tError 包__")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("类似 Python try ")
			fmt.Println("错误信息打印 -> ", err)
			fmt.Println("发送邮件给admin@sohu.com~")
		}
	}()
	num1, num2 := 10, 0
	res := num1 / num2
	fmt.Println("整数除以 零 ->", res)

}

func Run_Error() {
	Error_Demo()
}
