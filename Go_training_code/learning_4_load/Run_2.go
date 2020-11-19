package main

import (
	"fmt"
	"time"
)

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
