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
	Chan_Demo.Chan_Demo_8()
}
