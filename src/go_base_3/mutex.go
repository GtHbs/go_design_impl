package go_base_3

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value       int64
	normalValue int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		total.Lock() // 使用锁
		total.normalValue += i
		total.Unlock() // 释放锁
	}
}

func workerAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&total.value, int64(i)) // 使用原子类
	}
}

func Entry() {
	var wg sync.WaitGroup // 初始化信号量
	wg.Add(2)             // 新增两个信号量
	go worker(&wg)
	go worker(&wg)
	wg.Wait() // 等待信号量完成
	//fmt.Println(total.value)
	fmt.Println(total.normalValue)
}
