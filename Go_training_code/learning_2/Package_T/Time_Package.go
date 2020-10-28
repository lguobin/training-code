// Time_Package
package Package_T

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Time_1() {
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) + 1
		count++
		if n == 9 {
			break
		}
	}
	fmt.Println("生成 99 一共使用了 ", count)
}

func test_2() string {
	str, temp := "hello ", 0
	for i := 0; i < 100; i++ {
		str += strconv.Itoa(i)
	}
	return str + strconv.Itoa(temp)
}

func test_3() {
	fmt.Println("\t\t看看日期和时间相关函数和方法使用")
	now := time.Now()
	fmt.Printf("now = %v now_type = %T\n", now, now)

	fmt.Println("\t\t通过now可以获取到年月日，时分秒")
	fmt.Printf("年 = %v \n", now.Year())
	fmt.Printf("月 = %v \n", now.Month())
	fmt.Printf("月 = %v \n", int(now.Month()))
	fmt.Printf("日 = %v \n", now.Day())
	fmt.Printf("时 = %v \n", now.Hour())
	fmt.Printf("分 = %v \n", now.Minute())
	fmt.Printf("秒 = %v \n", now.Second())

	fmt.Println("\t\t格式化日期时间")
	fmt.Printf("当前年月日[ %d-%d-%d %d:%d:%d ]\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日转成 string 格式[ %d-%d-%d %d:%d:%d ]",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr = %v | %T\n", dateStr, dateStr)

	fmt.Println("\t\t格式化日期时间的第二种方式")
	fmt.Printf("打印年月日时分秒: \t%s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("打印年月日:    \t%s\n", now.Format("2006-01-02"))
	fmt.Printf("打印时分秒:    \t%s\n", now.Format("15:04:05"))
	fmt.Printf("打印年份:      \t%s\n\n", now.Format("2006"))
}

func test_4() {
	fmt.Println(`	需求，每隔1秒中打印一个数字，打印到100时就退出
	需求2: 每隔0.1秒中打印一个数字，打印到100时就退出`)
	i, start_time := 0, time.Now().Unix()
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i == 3 {
			break
		}
	}
	end_tiem := time.Now().Unix()
	fmt.Printf("Sleep Time总耗时: %v 秒\n", end_tiem-start_time)
}

func test_5() {
	fmt.Println("\t\t打印时间单位...")
	fmt.Println(time.Second)
	fmt.Println(time.Second * 60)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Millisecond * 1000)
	fmt.Println(time.Millisecond * 60000)

	now := time.Now()
	// Unix 和 UnixNano 的使用
	fmt.Printf("Unix时间戳 = %v、长度: %v| UnixNano时间戳 = %v、长度: %v\n",
		now.Unix(), len(fmt.Sprint(now.Unix())), now.UnixNano(), len(fmt.Sprint(now.UnixNano())))
}

func Time_Package() {
	fmt.Println("\t --- 调用 Time 包. --- ")
	Time_1()

	start_time := time.Now().Unix()
	test_2()
	end_time := time.Now().Unix()
	fmt.Printf("执行 test_2() 耗费时间为 %v 秒\n", end_time-start_time)

	test_3()
	test_4()
	test_5()

	a := 1603858268
	b, c := fmt.Sprint(a), strconv.Itoa(a)
	fmt.Printf("%T | %T\n", b, c)
	fmt.Println(b == c)

}
