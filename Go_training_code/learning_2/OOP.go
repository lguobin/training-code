// OOP
package main

import (
	"fmt"
	"learning_2/OOP"
)

func OOP__init__() {
	fmt.Println("\t面向对象 | 封装、继承、多态")
}

// 继承
type Father struct {
	Name string
	Age  int
}
type Son struct {
	Father
}

func (f *Father) run() {
	fmt.Println(f.Name+" running...", f.Age, "岁")
}

// 继承 - End

// 多重继承

type Father_B struct {
	Like string
}
type Son1 struct {
	Father
	Father_B
}
type Son2 struct {
	*Father
	*Father_B
}

// 多重继承 - End

// 多态
type BaseIntf interface {
	Process() string
}

func (temp *Father) Process() string {
	return "Process Father - Name: " + temp.Name
}
func (temp *Son) Process() string {
	return "Process Son|Father - Name: " + temp.Name
}
func PersonInfo(temp BaseIntf) {
	fmt.Println("多态", temp.Process())
}

// 多态 - End

func Bsisc_OOP() {
	OOP__init__()
	a := person.NewPerson("Tom")
	a.SetAge(18)
	fmt.Println(a)

	b := person.NewPerson("Mon")
	b.SetAge(151)
	fmt.Println(b)

	var d Son
	d.Father.Name = "张三"
	d.Father.Age = 31
	d.Father.run()

	//上述可以简写为：
	d.Name = "张三 - 儿子"
	d.Age = 15
	d.run()

	// 多重继承
	s1 := &Son1{Father{Name: "老A", Age: 30}, Father_B{Like: "设计师"}}
	fmt.Println("\t\t --- 多重继承 ---\n\t\t", s1)

	s2 := &Son2{&Father{"老B", 33}, &Father_B{"码农"}}
	fmt.Println("\t\t --- 多重继承 ---\n\t\t", s2.Father)

	// 多态与接口（interface）有关联
	PersonInfo(&Father{Name: "处理 - Person - Name"})
	PersonInfo(&Son{Father{Name: "处理 - Son|Father - Name"}})
}
