package Struct

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) test() float64 {
	return r.width * r.height
}

func (c Circle) test() float64 {
	return c.radius * c.radius
}

func Struct_Demo_2() {

	r1 := Rectangle{3, 6}
	r2 := Rectangle{3.3, 6.6}

	c1 := Circle{3.14}
	c2 := Circle{6.14}

	fmt.Printf("Rectangle 计算结果: -> %f | %.3f\n", r1.test(), r2.test())
	fmt.Printf("Circle 计算结果: -> %f | %f\n", c1.test(), c2.test())
}
