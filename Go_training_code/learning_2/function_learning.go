// slice_learning
package main

import (
	"bytes"
	"fmt"
)

func func__init__() {
	fmt.Println("\t学习信息 - func(函数)")
}

func joinStr(list ...string) string {
	var temp bytes.Buffer
	for _, s := range list {
		temp.WriteString(s)
	}
	fmt.Println(temp)
	return temp.String()
}

func Bsisc_func() {
	func__init__()
	a := []string{"AAA", "BBB", "CCC"}
	joinStr("pig", " and", " bird")
	joinStr(a...)
}
