package Struct

import "fmt"

type _Human interface {
	Len()
}
type _Student interface {
	_Human
}

type Test struct {
}

func (h *Test) Len() {
	fmt.Println("嵌入interface - 成功")
}

func Struct_Demo_4() {
	var s _Student
	s = new(Test)
	s.Len()
}
