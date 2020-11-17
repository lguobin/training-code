package main

import (
	"fmt"
	Math "learning_3_goroutine/Calculation"
	Chan_Demo "learning_3_goroutine/Channel"
	runtime "learning_3_goroutine/Goroutine"
	Req "learning_3_goroutine/Http"
)

func Go() {
	fmt.Println("\t --- Goroutine - 练习 --- ")
	runtime.Demo_0()
	runtime.Demo_1()
	runtime.Demo_2()
	Math.Math()
	runtime.Demo_3()
	runtime.Demo_4()
	runtime.Demo_5()

	// 多协程
	runtime.Demo_6()
	runtime.Demo_7()
	runtime.Demo_8()
	runtime.Demo_9()
}

func Lock() {
	runtime.Lock_UnLock()
	runtime.Read_write_lock()
	runtime.Wait_group_test()
	runtime.Map_safe_test()
}

func Channel_Demo() {
	Chan_Demo.Chan_Demo_0()
	Chan_Demo.Chan_Demo_1()
	Chan_Demo.Chan_Demo_2()
	Chan_Demo.Chan_Demo_3()
	Chan_Demo.Chan_Demo_4()
	Chan_Demo.Chan_Demo_5()

	// 生产者消费者模型
	Chan_Demo.Chan_Demo_6()
	Chan_Demo.Chan_Demo_7()
	Chan_Demo.Chan_Demo_8()
}

func Request() {
	Req.Request()
}
func main() {
	// Go()

	Lock()
}
