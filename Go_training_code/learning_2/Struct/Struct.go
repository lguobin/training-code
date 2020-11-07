// Struct
package Struct

import (
	"fmt"
)

func Struct__init__() {
	fmt.Println("\t结构 - 模拟构造函数")
	fmt.Println(`"贴士：因为Go没有函数重载，为了避免函数名字冲突，使用了
		'NewPerson'和'NewJob'两个不同的函数表示不同的'Person'构造过程。"`)
}

type Person struct {
	Name string
	Job  string
}

//构造父类
func NewPerson(_name, _job string) *Person {
	return &Person{Name: _name, Job: _job}
}

// 构造子类
type Student struct {
	Person
	level string
}

func NewStudent(_level string) *Student {
	p := &Student{}
	p.level = _level
	return p
}

// 自定义类型
type Interger int

func (i Interger) Less(j Interger) bool {
	fmt.Println("比较大小")
	return i < j
}

// 方法与函数 - 实现构造函数
// 一个 run_A 函数
func run_A(p *Person, temp string) {
	p.Name = temp
	fmt.Println("函数 run...", p.Name)
}

// 一个 run_B 方法
func (p *Person) run_B() {
	fmt.Println("方法 run...", p.Name)
}

// 定义一个表示点的结构体
type Point struct {
	X, Y int
}

// 非指针接收器
func (p_temp Point) Add(otherp Point) Point {
	return Point{p_temp.X + otherp.X, p_temp.Y + otherp.Y}
}

func Bsisc_Struct() {
	Struct__init__()

	name_Ptr_ := NewPerson("张三", "瞎扯蛋")
	fmt.Println(name_Ptr_)

	_student := NewStudent("第 三 等级")
	fmt.Println(_student)

	var temp Interger = 6
	fmt.Println(temp.Less(5))

	p_temp := &Person{}
	run_A(p_temp, "李四_AAAAAAA")

	T := &Person{"BBBBBB", "瞎扯蛋_BBBB"}
	T.run_B()

	p_1 := Point{1, 1}
	p_2 := Point{2, 2}
	result := p_1.Add(p_2)
	fmt.Println(result)
}
