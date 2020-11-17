package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Mall_createOrder() {
	// temp_status := fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	url := "http://test.sichuananpeng.com/mallActivity/createOrder"
	requestBody := "ext=&itemNum=1&request_time=&address=龙口东路333号&payType=2&shippingName=章鱼烧&cityName=广州市天河区博物馆奇妙夜&freight=0.00&actId=119&shippingPhone=13250111111&addressId=17"
	var jsonStr = []byte(requestBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("version", "1.9.36")
	req.Header.Add("platform", "android")
	req.Header.Add("Authorization", "")
	req.Header.Add("nonce", "")
	req.Header.Add("sign", "")
	req.Header.Add("timestamp", "")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// saveResponse(resp.Status, string(body))
	if resp.StatusCode != 200 {
		fmt.Println("response Status: ", resp.StatusCode)
		return
	}
}

func saveResponse(Statuscode, text string) {
	t := time.Now()
	msg := t.Format("2006-01-02 15:04:05") + " | " + Statuscode + " | " + text + "\n"
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
}

// Golang 实现线程池
type Pool struct {
	Queue         chan func() error
	RuntineNumber int
	Total         int

	Result         chan error
	FinishCallback func()
}

//初始化
func (self *Pool) Init(runtineNumber int, total int) {
	self.RuntineNumber = runtineNumber
	self.Total = total
	self.Queue = make(chan func() error, total)
	self.Result = make(chan error, total)
}

func (self *Pool) Start() {
	//开启 number 个goruntine
	for i := 0; i < self.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task()
				self.Result <- err
			}
		}()
	}

	//获取每个任务的处理结果
	for j := 0; j < self.RuntineNumber; j++ {
		res, ok := <-self.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	//结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback()
	}
}

//关闭
func (self *Pool) Stop() {
	close(self.Queue)
	close(self.Result)
}

func (self *Pool) AddTask(task func() error) {
	self.Queue <- task
}

func (self *Pool) SetFinishCallback(fun func()) {
	self.FinishCallback = fun
}

func String() []string {
	var a []string
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprint(i))
	}
	return a
}

func Run_B() {
	startTime := time.Now()
	var p Pool
	url := String()
	p.Init(9, len(url))
	for i := range url {
		u := url[i]
		p.AddTask(func() error {
			return Download(u)
		})
	}
	p.SetFinishCallback(DownloadFinish)
	p.Start()
	p.Stop()
	fmt.Printf("总耗时: %s\n", time.Since(startTime))
}

func Download(url string) error {
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("Download " + url)
	return nil
}

func DownloadFinish() {
	fmt.Println("Download finsh")
}
