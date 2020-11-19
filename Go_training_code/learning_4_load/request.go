package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

func Mall_createOrder() {
	// temp_status := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	url := "http://test.sichuananpeng.com/mallActivity/createOrder"
	requestBody := "ext=&itemNum=1&request_time=&address=龙口东路333号&payType=2&shippingName=章鱼烧&cityName=广州市天河区博物馆奇妙夜&freight=0.00&actId=119&shippingPhone=13250111111&addressId=17"
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Add("Authorization", "ZpIeMI6t9KfC8gkJ1/DkvRkwKCCqI8oGZawR6mS+w6CTXH1Vsb4tx8dU/ggGQx7vFG8qpZbsbhZB+CQyC2oBjI8f3iDZWMwhWw99ueVvAS+tZvgTd73K8LJ7YY76tDJZX1btK22mAbjxLQERfvviAv4Vd9CrsfihT5qNt9mZirAfPKeJEpkW+E8fdGK+NoRgxeEJrwehSr2M2UtsZu3SNsaudAh6Oz9Gv2KImNhCaegIOrTuYp4/QEBb1UGlm3YbPGcjod3drUXPo6Mp8Y816g==")
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add("version", "1.9.36")
	// req.Header.Add("platform", "android")
	// req.Header.Add("nonce", "f58f94ca3aeedf3dff65fc7e724f5efa")
	// req.Header.Add("sign", "54e6972e05b4b8234e8f34e8a5513091")
	// req.Header.Add("timestamp", "1605249282752")

	req.Header.Add("Authorization", "xfKCPtq6USOBtM7D4E4tCjNCHJwTWT5b6mIoXpCNnJNmPAFK/jANh/+98L9UVkZJW7jAs47Lv/d0qPZqyLNlSFerCFqahuqHcru8RckD5Dsq0iAP4zOKWMqvfjdBkTQ8eAZI50FjO53OkoYh8ZrIuBZ+lrq6VBPTCRTJySiKQv33DPcGIpsrkg3yuViSnhKQehY6zakKKUQ/vjGlTXTuIFqVCYxcWfSoZrd/xqOsUEJVQWvpX8hFV9eVTR3Ccs0tIOfAKqz9X0un8SQgaSpeeQ==")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("nonce", "4d6eaeec011e6def30b09a431b8bbd37")
	req.Header.Add("sign", "7f232b9f5107966424c7d1b6752996d2")
	req.Header.Add("timestamp", "1605694473075")
	req.Header.Add("version", "1.9.36")
	req.Header.Add("platform", "android")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	saveResponse(resp.Status, string(body))
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("response Status: ", resp.StatusCode)
		return
	}
}

func saveResponse(Statuscode, text string) {
	// t := time.Now()
	// msg := t.Format("2006-01-02 15:04:05") + " | " + Statuscode + " | " + text + "\n"
	msg := text + "\n"
	dir, _ := os.Getwd()
	filename := dir + "\\temp.txt"
	file, err := os.Open(filename)
	defer func() { file.Close() }()
	if err != nil && os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			fmt.Println("Failed to create file.")
		}
	}
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	defer f.Close()
	buf_write := bufio.NewWriter(f)
	buf_write.WriteString(msg)
	err = buf_write.Flush()
	if err != nil {
		fmt.Println("flush error : ", err)
	}
	// fmt.Println("文件保存在 ---> ", filename)
}

// var limit = 1000000000
var limit = 4
var wg sync.WaitGroup
var matex sync.Mutex

func task() {
	defer wg.Done()
	for {
		matex.Lock()
		if limit > 0 {
			Mall_createOrder()
			// saveResponse("test", fmt.Sprint(limit))
			limit--
		} else {
			fmt.Println("完成任务...")
			matex.Unlock()
			break
		}
		matex.Unlock()
	}
}

func run() {
	// startTime := time.Now().Unix()
	startTime := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(4)
	go task()
	go task()
	go task()
	go task()
	wg.Wait()
	fmt.Printf("总耗时: %s\n", time.Since(startTime))
}
