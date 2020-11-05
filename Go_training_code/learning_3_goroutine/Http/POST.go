package Http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Try_POST(url, contentType, body string) string {
	fmt.Println("\t --- 模拟发送 POST 请求，并返回响应体 --- ")
	req, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		fmt.Println("请求错误", err)
	}
	defer req.Body.Close()
	response, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("请求错误", err)
	}
	fmt.Printf("POST | 响应status: %v\n", string(req.Status))
	// fmt.Printf("响应体: %v\n", string(response))
	return string(response)
}
