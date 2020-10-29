package main

import (
	Algorithm "algorithm/Sort"
)

func main() {
	// 测试一把
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	Algorithm.BinaryFind(&arr, 0, len(arr)-1, -6)
	Algorithm.BubbleSort(&[5]int{101, 18, 10, 89, 1000})
}
