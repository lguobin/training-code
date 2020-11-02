package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Slice 转 String
func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

// String 转 Slice
func Slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func string_to_slice() {
	fmt.Println("\t --- 字符串转切片 ---")
	b := []byte("hello world")
	a := String(b)
	b[0] = 'a'
	println(a) //output  aello world

	aa := "hello world"
	temp := Slice(aa)
	temp[0] = 'a' // 这里就等着崩溃吧
	fmt.Println(temp[0])
}
