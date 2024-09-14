package go_base_3

import (
	"sync/atomic"
	"time"
)

var config atomic.Value

func loadConfig() {}

func AtomicEntry() {
	config.Store(loadConfig)
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig)
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {

		}()
	}
}
