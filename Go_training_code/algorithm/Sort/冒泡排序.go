package Sort

import "fmt"

func BubbleSort(arr *[5]int) {

	fmt.Println("\t冒泡排序")
	fmt.Println("排序前arr=", (*arr))
	// 冒泡排序..一步一步推导出来的
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交换
				(*arr)[j+1], (*arr)[j] = (*arr)[j], (*arr)[j+1]
			}
		}

	}
	fmt.Println("排序后arr=", (*arr))
}
