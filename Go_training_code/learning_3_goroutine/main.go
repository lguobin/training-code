package main

import (
	"fmt"
	Math "learning_3_goroutine/Calculation"
	Chan_Demo "learning_3_goroutine/Channel"
	Mode "learning_3_goroutine/Goroutine"
)

func Go() {
	fmt.Println("\t --- Goroutine - 练习 --- ")
	Mode.Demo_0()
	Mode.Demo_1()
	Mode.Demo_2()
	Math.Math()
	Mode.Demo_3()

}
func Channel_Demo() {
	Chan_Demo.Chan_Demo_1()
	Chan_Demo.Chan_Demo_2()
}
func main() {
	// Go()

	Chan_Demo.Chan_Demo_1()
	Chan_Demo.Chan_Demo_2()
}
