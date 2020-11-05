package Http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("请求错误", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("请求错误", err)
	}
	fmt.Println("request status", resp.Status)
	fmt.Println("request Headers", resp.Header)
	fmt.Println("request done", string(body))
}

// 编码类型为：application/json
func httpPost(requestBody, url, USER_UID, USER_IS_LOGIN string) {
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("USER_UID", USER_UID)
	req.Header.Add("USER_IS_LOGIN", USER_IS_LOGIN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
}

// 编码为 application/x-www-form-urlencoded
func httpDo() {
	client := &http.Client{}
	const uriBase = "http://10.84.135.139:8089"
	const uriPath = "/inter/xpsfsdfsdgan/inter?method=batchcashlist"
	const url = uriBase + uriPath
	req, err := http.NewRequest("POST", url, strings.NewReader("pids=1829587349042601,985162418600072,985162418617402"))
	if err != nil {
		// handle error
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("USER_UID", "4093715774")
	req.Header.Add("USER_IS_LOGIN", "1")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

// 并发测试方法
func testBingfa(n int) {
	requestBody := `{'key1':'value1','key2':'value2'}`
	url := "http://httpbin.org/post"
	USER_UID := "2222342870"
	USER_IS_LOGIN := "1"
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			httpPost(requestBody, url, USER_UID, USER_IS_LOGIN) // 被测试方法
			wg.Done()
		}()
	}

	go func() {
		// 开一个http端口来侦听，比如
		fmt.Println("开一个http端口来侦听，比如: pprof start...")
		// fmt.Println(http.ListenAndServe(":9876", nil))
		fmt.Println("当前开启goruntime数量: ", runtime.NumGoroutine())
	}()

	wg.Wait()
	fmt.Println("ending----")
}

func Request() {
	// httpGet(url)
	// testBingfa(3)

	// url := "http://httpbin.org/"
	// Try_GET(url + "get")
	// contentType := "application/json"
	// json_POST := `{'key1':'value1','key2':'value2'}`
	// Try_POST(url+"post", contentType, json_POST)

	Run_time()
}
