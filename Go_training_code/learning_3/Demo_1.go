package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello...")
}

type Manager struct {
	// 匿名结构体
	User
	title string
}

func Info(o interface{}) {

	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: \t")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %s = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("方法名: %6s: %s\n", method.Name, method.Type)
	}
}

func Demo_1() {
	fmt.Println("结构体反射..")
	u := User{1, "OK", 12}
	Info(u)

	m := Manager{User: User{2, "No", 24}, title: "测试"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}
