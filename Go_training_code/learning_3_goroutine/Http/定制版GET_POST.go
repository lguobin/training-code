package Http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func ParseResponse(response *http.Response) (map[string]interface{}, error) {
// 	var result map[string]interface{}
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err == nil {
// 		err = json.Unmarshal(body, &result)
// 	}
// 	return result, err
// }

// 多种请求方式, 修改" Content-Type "即可
func Requests_MOD(method, url, requestBody, ContentType, token string) {

	fmt.Println("\t --- 请求开始 --- ")
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", token)
	req.Header.Add("Accept", ContentType)
	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("response Status: %d\nresponse Body: %s\n", resp.StatusCode, string(body))
	defer resp.Body.Close()
}

func Run_time() {
	method := "POST"
	url := "http://127.0.0.1/login/"
	requestBody := `{"name":"15920150690","password":"saiyao1234"}`
	ContentType := "application/json"
	token := `JWT eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvaWQiOiJlNTQ5YjQzYmZjZjQ0YWFlOWM2NDUyMWU5NjE0ZDBiOCIsIm1hc2siOiJmZmZmZmZmZmZmZmZmZmZmZmZmZmZmZmZmIiwicm9sIjoib3BlcmF0b3IiLCJleHAiOjE2MDUwNzg3NzMsImxpYyI6NH0.J_WqWAtJYbbuhEwHbpDCUwuzIkkRanbpunNYRjKjORY`
	Requests_MOD(method, url, requestBody, ContentType, token)
}
