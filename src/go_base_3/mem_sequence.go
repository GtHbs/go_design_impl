package go_base_3

import (
	"fmt"
	"sync"
)

func MemSequenceWithChan() {
	ch := make(chan int)
	var s string
	go func() {
		s = "hello world"
		ch <- 1
	}()
	<-ch
	fmt.Println(s)
}

func MemSequenceWithMutex() {
	var mutex sync.Mutex
	mutex.Lock()
	var s string
	go func() {
		s = "hello world"
		mutex.Unlock()
	}()
	mutex.Lock()
	fmt.Println(s)
}

func MultiChan() {
	limit := make(chan int, 10)
	for i := 0; i < cap(limit); i++ {
		go func() {
			fmt.Println("lazy")
			limit <- 1
		}()
	}
	for i := 0; i < cap(limit); i++ {
		<-limit
	}
}
