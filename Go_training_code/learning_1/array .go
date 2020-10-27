// 序列化
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal_temp struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	petName string `json:"pet_Name"`
}

func Run_Animal() {
	animals := []Animal_temp{
		Animal_temp{"dog", 3, "旺旺..."},
		Animal_temp{"cat", 2, "喵喵..."},
	}
	bs, err := json.Marshal(animals)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("序列化打印 ->  : %v\n", string(bs))
}
