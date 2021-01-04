package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0] == b[0])
	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	var times [5][0]int // `[0]int`大小是0，因此整个数组占用的内存大小依然是0
	for range times {
		fmt.Println("hello")
	}

	// 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	fmt.Println(s1, s2, s3)

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1, line2, line3)

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Println(decoder1, decoder2)

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
	fmt.Println(unknown1, unknown2)

	// 管道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList)

	// 长度为0的数组在内存中并不占用空间, 管道的同步操作
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	// 我们并不关心管道中传输数据的真实类型, 无类型的匿名结构体
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2

	// Go语言字符串的底层结构在 `reflect.StringHeader` 中定义
	type StringHeader struct {
		Data uintptr
		Len  int
	}

	var data = [...]byte{
		'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',
	}
	fmt.Println(data)

	//
	s := "hello, world"
	// hello := s[:5]
	// world := s[7:]
	// s1 = "hello, world"[:5]
	// s2 = "hello, world"[7:]
	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

	_byt := []byte("Hello, 世界")
	fmt.Printf("%#v\n", string(_byt))
	fmt.Printf("%#v\n", _byt)
	fmt.Println("\xe4\xb8\x96")

	// for i, c := range []byte("世界abc") {
	// 	fmt.Println(i, c)
	// }

	// s = "\xe4\x00\x00\xe7\x95\x8cabc"
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%d %x\n", i, s[i])
	// }

	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界

	// for range 对字符串的迭代模拟实现
	// s = "AB"
	// forOnString(s)

	// 模拟字符串 string 转 []byte
	s = "ABCDEF"
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		// c := s[i]
		// p[i] = c
		p[i] = s[i]
	}
	fmt.Println(s)
	fmt.Println(p)

	// 模拟 []byte 转 string
	s_byte := []byte{65, 66, 67, 68, 69, 70}
	bytes2str(s_byte)

	// 模拟 []rune 转 string
	str2runes(s)

}

// ------------------------
// ------------------------
func forOnString(s string, forBody func(i int, r rune)) {
	for i := 0; len(s) > 0; {
		r, size := utf8.DecodeRuneInString(s)
		forBody(i, r)
		s = s[size:]
		i += size
	}
}
func bytes2str(s_byte []byte) (ptr string) {
	// 模拟 []byte 转 string
	s_data := make([]byte, len(s_byte))
	for i, d := range s_byte {
		s_data[i] = d
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&ptr))
	hdr.Data = uintptr(unsafe.Pointer(&s_data[0]))
	hdr.Len = len(s_byte)
	fmt.Println(ptr)
	return ptr
}
func str2runes(s string) []rune {
	// 模拟 []rune 转 string
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	fmt.Println(p)
	return []rune(p)
}
