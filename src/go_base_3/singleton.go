package go_base_3

import (
	"sync"
	"sync/atomic"
)

type Singleton struct{}

type Once struct {
	mu   sync.Mutex
	done int32
}

var (
	instance    *Singleton
	initialized int32
	mu          sync.Mutex
)

func Instance() *Singleton {
	if atomic.LoadInt32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		atomic.StoreInt32(&initialized, 1)
		instance = &Singleton{}
	}
	return instance
}

func (o *Once) Do(f func()) {
	if atomic.LoadInt32(&o.done) == 1 {
		return
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		defer atomic.StoreInt32(&o.done, 1)
		f()
	}
}

var (
	instances *Singleton
	once      sync.Once
)

func InstanceOnce() *Singleton {
	once.Do(func() {
		instances = &Singleton{}
	})
	return instances
}
