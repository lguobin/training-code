// Time_Package
package Package_T

import (
	"fmt"
	"math/rand"
	"time"
)

func Time_Package() {
	fmt.Println("调用 Time 包.")

	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) + 1
		count++
		if n == 9 {
			break
		}
	}
	fmt.Println("生成 99 一共使用了 ", count)
}
