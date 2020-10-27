// Loin_check
package Package_T

import (
	"fmt"
)

func Login_check() {
	login_count := 3
	username, password := "admin", "admin"
	var n_temp, p_temp string
	fmt.Printf("Login_check, 登录信息: %v | %v\n", username, password)

	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		// fmt.Scanln(&n_temp)
		fmt.Println("请输入密码")
		// fmt.Scanln(&p_temp)

		if n_temp == username && p_temp == password {
			fmt.Println("登录成功...")
			break
		} else {
			login_count--
			fmt.Printf("用户名 || 密码错误, 还剩 %v 次登录机会.\n", login_count)
		}
		if login_count <= 1 {
			// 最后一次循环直接给与真实 用户名 & 密码
			n_temp, p_temp = "admin", "admin"
		}
	}
	if login_count == 0 {
		fmt.Println("登录次数使用完，请明天再来重试!")
	}
}

func Count_Num_Test() {
	var Positive_Count, Negative_Count, Number int
	var Pos_list, Neg_list []int // 初始化 list
	for {
		fmt.Println("请输入一个整型, 退出输入: 0")
		fmt.Scanln(&Number)

		if Number == 0 {
			break
		}
		if Number > 0 {
			Positive_Count++
			Pos_list = append(Pos_list, Number) // list 追加元素
			continue
		} else {
			Negative_Count++
			Neg_list = append(Neg_list, Number) // list 追加元素
			continue
		}
	}
	fmt.Printf("正数个数是 %v ,负数的个数是 %v\n", Positive_Count, Negative_Count)
	fmt.Printf("正数列表：%v\n负数列表：%v\n", Pos_list, Neg_list)
}

func Checker_Sum() {
	var a, b int
	a, b = 50, 51
	if a+b >= 100 {
		fmt.Println("两数之和大于100,", a+b)
	} else {
		fmt.Println("两数之和小于100,", a+b)
	}
}

func Calc() {
	var a, b float64 = 2.2, 3.401
	var operator byte = '+'
	var result float64

	switch operator {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		fmt.Println("操作符错误...")
	}
	fmt.Println("结果: ", result)
}

func Recursive_call(temp int) {
	if temp > 2 {
		temp--
		Recursive_call(temp)
	}
	fmt.Println(temp)
}

func bulidinfunction() {
	num1 := 100
	fmt.Printf("num1的类型%T , 值=%v , 内存地址=%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	fmt.Printf("num2的类型%T , 值=%v , 内存地址=%v, 指针值=%v\n", num2, num2, &num2, *num2)
}

func digui(n int) int {
	fmt.Println(`
	解题:
		1. 假设 n = 6
		2. 第 1 轮递归: digui( 6 - 1) * 6 ==>> 6 * 5
		3. 第 2 轮递归: digui( 5 - 1) * 6 ==>> 6 * 5 * 4
		4. 第 3 轮递归: digui( 4 - 1) * 6 ==>> 6 * 5 * 4 * 3
		5. 第 4 轮递归: digui( 3 - 1) * 6 ==>> 6 * 5 * 4 * 3 * 2
		6. 第 5 轮递归: digui( 2 - 1) * 6 ==>> 6 * 5 * 4 * 3 * 2 * 1
		7. 第 6 轮递归: digui( 1 - 1) * 6 ==>> 6 * 5 * 4 * 3 * 2 * 1 * 0 | 已超过 n 最小限制，直接跳出递归
		total: 6 *5 *4 *3 *2 *1 = 720
	`)
	if n < 2 {
		return n
	}
	_temp := digui(n-1) * n
	fmt.Printf("n -> %v\t结果:%v\n", n, _temp)
	return _temp
}

func fib_function(n int) int {
	if n < 2 {
		return n
	}
	return fib_function(n-1) + fib_function(n-2)

}

func f(n int) int {
	// 题2：求函数值已知 f(1)=3; f(n) = 2*f(n-1)+1; 请使用递归的思想编程，求出 f(n)的值?
	if n == 1 {
		return 3
	}
	_temp := 2*f(n-1) + 1
	fmt.Printf("n -> %v\t结果:%v\n", n, _temp)
	return _temp
}

func Run() {
	Recursive_call(10)
	bulidinfunction()

	// 递归 测试一下
	fmt.Println(digui(6))

	// fmt.Println(fib_function(10)) // 55

	// fmt.Println("f(1)=", f(1)) // n -> 1	结果:3
	// fmt.Println("f(5)=", f(4)) // n -> 4	结果:31
	// fmt.Println("f(6)=", f(6)) // n -> 6	结果:127
}
