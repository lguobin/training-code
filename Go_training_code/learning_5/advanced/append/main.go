package main

import "fmt"

func main() {
	const N = 1024
	var a [N]int

	// appned 触发扩容
	// fmt.Println(a[: N-1 : N])
	x := append(a[:N-1:N], 0, 9)
	y := append(a[:N:N], 9)
	fmt.Println(cap(x), cap(y))
}
