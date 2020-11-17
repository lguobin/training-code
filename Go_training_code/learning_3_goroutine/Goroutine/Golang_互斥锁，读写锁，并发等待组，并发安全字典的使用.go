package Goroutine

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// sync同步包
// 1、互斥锁
func Lock_UnLock() {
	var lock sync.Mutex
	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("V1 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("V1 free lock at " + time.Now().String())
	}()
	time.Sleep(time.Second / 10)

	go func() {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("V2 get lock at " + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("V2 free lock at " + time.Now().String())
	}()
	// 等待所有线程执行完毕(让主线程不会死)
	time.Sleep(time.Second * 4)
}

// 2、读写锁
var rwLOCK sync.RWMutex

func Read_write_lock() {
	// 获取读锁
	for i := 0; i < 5; i++ {
		go func(val int) {
			rwLOCK.RLocker()
			// 注意这里读锁的获取和释放是一样的，我有点不解
			defer rwLOCK.RLocker()
			fmt.Println("read lock" + strconv.Itoa(val) + "get rlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second / 10)

	// 获取写锁
	for i := 0; i < 5; i++ {
		go func(val int) {
			// 这里和互斥锁是相同的接口
			rwLOCK.Lock()
			defer rwLOCK.Unlock()
			fmt.Println("write lock" + strconv.Itoa(val) + "get wlock at " + time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second * 15)
}

// 并发等待组
func Wait_group_test() {
	var waitGroup sync.WaitGroup
	routineSize := 5
	waitGroup.Add(routineSize)
	for i := 0; i < 5; i++ {
		go func(val int) {
			fmt.Println("work" + strconv.Itoa(val) + " is done at" + time.Now().String())
			time.Sleep(time.Second)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	fmt.Printf("all work done at %s", time.Now().String())
}

// 并发安全字典
func Map_safe_test() {
	var syncMap sync.Map
	var waitGroup sync.WaitGroup
	routineSize := 5
	waitGroup.Add(routineSize)
	for i := 0; i < routineSize; i++ {
		go func(begin int) {
			for i := begin; i < begin+3; i++ {
				syncMap.Store(i, i)
			}
			//通知数据添加完毕
			waitGroup.Done()
		}(i * 10)
	}

	// 开始等待
	waitGroup.Wait()
	var size int
	syncMap.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	fmt.Printf("syncMap current size is %s\n", strconv.Itoa(size))
	val, ok := syncMap.Load(0)
	if ok {
		fmt.Printf("key 0 has value:%v\n", val)
	}

}
