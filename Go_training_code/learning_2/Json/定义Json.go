package Json

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"` // 反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"Birthday"`
	Attack   float64 `json:"Attack"`
}

func test_Json() {
	// 定义一个json
	temp_data := Monster{Name: "牛魔王", Age: 100, Birthday: "111111111", Attack: 106.66}
	fmt.Println(temp_data)

	data, err := json.Marshal(&temp_data)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("temp_data Json 序列化后 = %v\n", string(data))
}

func test_Map() {
	// 定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 12
	a["address"] = "北京. 北京. "
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("a Map 序列化后 = %v\n", string(data))
}

func test_Slice() {
	// 演示对切片进行序列化, 我们这个切片 []map[string]interface{}
	var slice []map[string]interface{}
	var temp map[string]interface{}
	temp = make(map[string]interface{})
	temp["name"] = "jack"
	temp["age"] = "7"
	temp["address"] = "北京"
	slice = append(slice, temp)

	var m2 map[string]interface{}
	// 使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("Slice 序列化后 = %v\n", string(data))
}

func test_Bsisc() {
	temp := "ABCDE"
	mp := make(map[int]interface{})
	for index, value := range temp {
		mp[index] = string(value)
	}
	fmt.Println("字符串加入 Map 打印", mp)

	//
	bbb := make(map[int]interface{})
	for i := 65; i <= 69; i++ {
		bbb[i] = string(i)
	}
	fmt.Println("整形加入 Map 打印", bbb)

	// 字符串序列化
	data, err := json.Marshal(temp)
	if err != nil {
		fmt.Printf("序列化错误 err = %v\n", err)
	}
	fmt.Printf("字符串 序列化后 = %v\n", string(data))
}

func Json_Demo() {
	fmt.Println("\t --- 定义 Json 格式 --- ")
	test_Json()
	test_Map()
	test_Slice()
	test_Bsisc()
}
