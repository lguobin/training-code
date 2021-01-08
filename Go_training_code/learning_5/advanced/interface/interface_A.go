package main

import "fmt"

type Tagget interface {
	// 必须与结构体方法返回参数一致
	AA() *mission
	BB(s string) *mission
}

type mission struct {
	Name  string
	Count int
}

func mission_a(s string) string {
	return s
}

func (this *mission) AA() *mission {
	return &mission{
		Name:  this.Name,
		Count: 111,
	}
}
func (this *mission) BB(s string) *mission {
	return &mission{
		Name:  s,
		Count: 2222,
	}
}

func InterFace_A() {
	// 接口使用方法 1
	temp := mission{Name: "接口使用方法 1"}
	var tagget Tagget = &temp
	a := tagget.AA()

	msg := "测试2"
	temp = mission{}
	// fmt.Println(temp.BB(msg))
	b := tagget.BB(msg)
	fmt.Println(a, b)

	// 接口使用方法 2
	var tagget_2 Tagget
	msg = "接口使用方法 2"
	tagget_2 = &mission{Name: msg}
	c := tagget_2.AA()
	d := tagget_2.BB(msg)
	fmt.Println(c, d)

	// 接口使用方法 3
	inter(temp)

	Myinter(1)
	Myinter("泛型接口")
	Myinter([]int{1, 2, 3})
	Myinter([...]int{1, 2, 3})
	Myinter(map[int]int{1: 1, 2: 2})
	Myinter(make(map[interface{}]interface{}))
}

func inter(this mission) {
	// 接口使用方式 3
	fmt.Println(this.AA(), this.BB("接口使用方法 3"))
}

func Myinter(a interface{}) {
	// 泛型
	fmt.Println(a)
}
