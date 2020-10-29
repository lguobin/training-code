package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	//判断users中是否有name
	//v , ok := users[name]
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name + "\n" //示意
	}
}

func Map_Demo() {
	// 第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	// 演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	// 当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	// 演示map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	// 如果希望一次性删除所有的key
	// 1. 遍历所有的key,如何逐一删除 [遍历]
	// 2. 直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

}

//定义一个学生结构体
type Stu struct {
	Name string
	Age  int
	City string
}

func Bsisc_Map() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println(users)

	Map_Demo()

	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 28, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println("学生列表: ", students)

	// 遍历学生列表
	for k, v := range students {
		fmt.Printf("学生的编号是: %v\n", k)
		fmt.Printf("学生的名字是: %v\n", v.Name)
		fmt.Printf("学生的年龄是: %v\n", v.Age)
		fmt.Printf("学生的地址是: %v\n", v.City)
		fmt.Println()
	}
}
