package go_base_3

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerFunc(cancel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-cancel:
			return
		default:
			fmt.Println("worker")
		}
	}
}

func cancelByContext(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			continue
		}
	}
}

func ExitMain() {
	wg := new(sync.WaitGroup)
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerFunc(cancel, wg)
	}
	time.Sleep(time.Second)
	cancel <- true
	wg.Wait()
}

func ExitByContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 当前routine最大执行时长为10s
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			err := cancelByContext(ctx, &wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	time.Sleep(time.Second * 20)
	cancel()
	wg.Wait()
}
