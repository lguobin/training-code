// struct
// 类似 class ，可面向对象
package main

import (
	"fmt"
)

type Animal struct {
	Name    string
	Age     int
	petName string // 私有成员
	// 大写表示这个结构体是公有的，小写表示结构体属于私有的，在其它地方不能使用
}

func (a Animal) Sleep() {
	fmt.Printf("『 %s 』 -> 正在睡觉...\n", a.Name)
}
func (a Animal) SetPetName(petName string) {
	a.petName = petName
}
func (a *Animal) Ptr_SetPetName(petName string) {
	a.petName = petName
	// NOTE: 这里的 a 是一个指针
	// NOTE: 以下这种方式也是可以的，go 如果碰到指针会自动帮我们处理，所以使用起来更方便
	// (*a).petName = petName
}

// 实现构造函数
func NewAnimal(name string, age int) *Animal {
	a := Animal{
		Name: name,
		Age:  age,
	}
	return &a
}

// 组合 vs 继承
type Dog struct {
	Animal
	Color string
}

func (d Dog) Sleep() {
	fmt.Println("Dog 输出自己的睡眠 -> Sleep")
}

func Running_Demo() {
	a := Animal{Name: "老王八の龟", Age: 1}
	a.Sleep()

	a.SetPetName("私有龟")
	fmt.Println(a.SetPetName)

	// 指针接收者(pointer receiver) vs 值接收者(value receiver)
	_Ptr := &Animal{Name: "老王八の龟_2", Age: 1}
	_Ptr.Ptr_SetPetName("私人成员")
	fmt.Println(_Ptr, _Ptr.Ptr_SetPetName)

	// 类似构造函数
	b := NewAnimal("凌家小猫", 18)
	fmt.Println(b)

	// 组合 vs 继承
	c := Dog{}
	c.Name = "独角兽"
	c.Sleep()
}

/////////////////////////////////

// 匿名结构
func niming_Structure_Demo() {
	a := &struct {
		Name string
		Age  int
	}{
		Name: "张三",
		Age:  18,
	}
	fmt.Printf("匿名结构 -> %v\n", a)

	a.Name = "李四"
	fmt.Printf("匿名结构-修改 -> %v\n", a)
}

// 结构方法
type _A struct {
	_B
	_C
}
type _B struct {
	Name string
}
type _C struct {
	Name string
}

func Struct_method() {
	a := _A{_B: _B{Name: "BBB"}, _C: _C{Name: "CCC"}}
	// fmt.Printf("结构体提升等级: %v\n", a.C.Name)
	fmt.Println(a._B.Name, a._C.Name)

}
