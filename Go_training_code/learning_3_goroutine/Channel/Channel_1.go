package Channel

import (
	"fmt"
	"time"
)

func Chan_Demo_1() {

	// 声明管道
	var write_intChan chan<- int // 声明只写管道
	write_intChan = make(chan int, 3)
	write_intChan <- 20
	write_intChan <- 30
	// write_temp := <-write_intChan
	// for v := range write_intChan {
	// 	fmt.Println("write_intChan ", v)
	// }
	// invalid operation: <-write_intChan (receive from send-only type chan<- int)

	var read_intChan <-chan int // 声明只读管道
	read_intChan = make(chan int, 3)
	// read_intChan <- 666
	// invalid operation: read_intChan <- 666 (send to receive-only type <-chan int)
	fmt.Println(read_intChan)

	// 结束声明

	// 1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值 = %v intChan本身的地址 = %p\n", intChan, &intChan)

	// 3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	// 如果从channel取出数据后，可以继续放入
	<-intChan
	// 注意点, 当我们给管写入数据时，不能超过其容量
	intChan <- 98

	// 4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	// 5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	// num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

	// 通过遍历获取数据
	intChan <- 1
	intChan <- 2
	intChan <- 3
	close(intChan)
	print("\t ---- 不锁管道会抛异常:\n\t[ fatal error: all goroutines are asleep - deadlock! ]\n\n")
	for v := range intChan {
		fmt.Println(" 管道遍历打印 = ", v)
	}
}

// -----------------------------
// -----------------------------

// 读取数据
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, err := <-intChan
		if !err {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

// 写入数据
func writeData(intChan chan int) {
	for i := 0; i <= 10; i++ {
		intChan <- i
		fmt.Println("writeData ", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func Chan_Demo_2() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go readData(intChan, exitChan)
	go writeData(intChan)

	// time.Sleep(time.Second * 10)
	for {
		_, err := <-exitChan
		if !err {
			break
		}
	}
}
