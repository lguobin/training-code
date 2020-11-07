package Struct

import (
	"fmt"
)

type Human struct {
	name      string
	phone_num string
}

type D_Student struct {
	Human  // 匿名字段
	S_aaaa string
}
type Employee struct {
	Human  // 匿名字段
	E_bbbb string
}

// Human实现 SayHi 方法
func (h Human) SayHi() {
	fmt.Printf("你好, 我是 %s | 手机 %s\n", h.name, h.phone_num)
}

// Human实现 song 方法
func (h Human) Song(song string) {
	fmt.Printf("唱歌: ---> %s\n\n", song)
}

// Interface Men被Human,D_Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Song(song string)
}

func Struct_Demo_3() {
	a := D_Student{Human{"张三", "13111111111"}, "AAAAAAAAAAA"}
	b := D_Student{Human{"李四", "13222222222"}, "BBBBBBBB"}
	c := Employee{Human{"王五", "1355555555555"}, "CCCCC"}
	d := Employee{Human{"赵六", "1326666666666"}, "DDDDDDDD"}

	// 定义Men类型的变量
	var temp Men
	temp = a
	fmt.Println("\t打印 [ A ]")
	temp.SayHi()
	temp.Song("汪峰 - 怒放的生命")

	temp = b
	fmt.Println("\t打印 [ B ]")
	temp.SayHi()
	temp.Song("汪峰 - 河流")

	temp = c
	fmt.Println("\t打印 [ C ]")
	temp.SayHi()
	temp.Song("SHE - 伏地魔")

	temp = d
	fmt.Println("\t打印 [ D ]")
	temp.SayHi()
	temp.Song("刘德华 - 小男孩")

	// 最快捷的方式
	// 通过 slice 配置循环
	fmt.Println("\t --- 通过 slice 配置循环 --- \n\n")
	var temp_slice []Men
	for i := 1; i < 5; i++ {
		x := fmt.Sprint(i)
		temp_slice = append(temp_slice, D_Student{Human{x + " - A", x}, "****"})
	}
	for _, va := range temp_slice {
		va.SayHi()
		va.Song("随机歌曲")
	}
}
