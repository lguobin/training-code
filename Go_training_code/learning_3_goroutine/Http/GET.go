package Http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Try_GET(url string) string {
	fmt.Println("\t --- 模拟发送 GET 请求，并返回响应体 --- ")
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("请求错误", err)
	}
	defer req.Body.Close()
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("请求错误", err)
	}
	fmt.Printf("GET | 响应status: %v\n", string(req.Status))
	// fmt.Printf("响应体: %v\n", string(response))
	return string(response)
}
