package Json

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster_2 struct {
	Name     string
	Age      int
	Birthday string
	Attack   float64
}

func string_struct() {
	str := "{\"Name\":\"牛魔王~~~\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var temp Monster_2

	err := json.Unmarshal([]byte(str), &temp)
	if err != nil {
		fmt.Printf("反序列错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster = %v monster.Name=%v \n", temp, temp.Name)

}

func Map_to_struct() {
	// 定义一个map
	a := make(map[string]interface{})
	a["name"] = "红孩儿~~~~~~"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	temp := string(data)
	var New_mp map[string]interface{}
	err_ := json.Unmarshal([]byte(temp), &New_mp)
	if err_ != nil {
		fmt.Printf("反序列错误 err=%v\n", err_)
	}
	fmt.Printf("Map\\json 反序列化后 New_mp = %v\n", New_mp)
}
func Slice_to_struct() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	fmt.Printf("slice 转换前类型: %T\n", str)

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("slice 反序列化后 slice=%v\n", slice)
}

func test() {
	Map__Josn_li := make(map[string]interface{})
	for i := 65; i <= 70; i++ {
		temp := string(i)
		Map__Josn_li[temp] = string(i + 32)
	}
	data, _ := json.Marshal(Map__Josn_li)
	temp := string(data)
	fmt.Printf("转换前: %v\t%T\n转换后: %v\t%T\n", Map__Josn_li, Map__Josn_li, temp, temp)
}

func Json_to_struct() {
	fmt.Println("\t --- 演示将json字符串，反序列化成 struct --- ")
	string_struct()
	Map_to_struct()
	Slice_to_struct()

	test()
}
