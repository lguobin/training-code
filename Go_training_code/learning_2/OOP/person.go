// OOP_Person.go
package person

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

// 工厂函数（类似构造函数）
func NewPerson(temp string) *person {
	return &person{Name: temp}
}

func (p *person) SetAge(temp int) {
	if temp > 0 && temp < 150 {
		p.Age = temp
	} else {
		fmt.Println("年纪错误..")
	}
}

func (p *person) GetAge() int {
	return p.Age
}
