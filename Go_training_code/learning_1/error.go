// error
package main

import (
	"errors"
	"fmt"
)

// defer 语句
func test_Defer_Demo() string {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("函数体 - 执行")
	return "最后返回的值"
}

// 定义错误
// Divide compute int a/b
func testError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为 -> 零")
	}
	return a / b, nil
}

// 自定义自己的业务异常
// type test_Error_Struct struct {
// 	Code    int32
// 	Message string
// }
// func (e *test_Error_Struct) string {
// 	return fmt.Sprintf("[test_Error_Struct] Code=%d, Message=%s", e.Code, e.Message)
// }
// func NewArticleError(code int32, message string) error {
// 	return &test_Error_Struct{
// 		Code:    code,
// 		Message: message,
// 	}
// }
