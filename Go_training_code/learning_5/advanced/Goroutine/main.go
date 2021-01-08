package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total_struct struct {
	sync.Mutex
	count int
}

func worker(wg *sync.WaitGroup) {
	// 为了保证`total_struct.value += i`的原子性
	// 通过`sync.Mutex`加锁和解锁来保证该语句在同一时刻只被一个线程访问
	defer wg.Done()
	for _i := 0; _i <= 10; _i++ {
		total_struct.Lock()
		total_struct.count += _i
		total_struct.Unlock()
	}
}

var total int64

func atomic_worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var _i int64
	for _i = 0; _i <= 10; _i++ {
		atomic.AddInt64(&total, _i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go worker(&wg)
	go worker(&wg)

	go atomic_worker(&wg)
	go atomic_worker(&wg)

	wg.Wait()
	fmt.Println(total_struct.count)
	fmt.Println(total)
}
