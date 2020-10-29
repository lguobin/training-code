package Sort

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	fmt.Println("\t二分查找")
	// 判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal {
		// 说明我们要查找的数，应该在  leftIndex --- middel-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 说明我们要查找的数，应该在  middel+1 --- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}
}
