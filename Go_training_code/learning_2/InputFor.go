// InputFor
package main

import (
	"fmt"
)

func Menu() {
	var temp int
	for {
		// 1. 先输出这个主菜单
		fmt.Println("-----------------家庭收支记账软件-----------------")
		// fmt.Println("			1 收支明细")
		// fmt.Println("			2 登记收入")
		// fmt.Println("			3 登记支出")
		// fmt.Println("			4 退出")
		// fmt.Print("请选择(1-4):")
		fmt.Scanln(&temp)

		fmt.Printf("请选择(1-4): %d\n", temp)

		if temp == 1 {
			break
		}
	}
}
