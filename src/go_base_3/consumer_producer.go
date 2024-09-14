package go_base_3

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(num int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * num
	}
}

func Consumer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func ConsumerEntry() {
	channel := make(chan int, 100)
	go Producer(1, channel)
	go Producer(2, channel)
	go Consumer(channel)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v) \n", <-sig)
}
