package go_base_3

import "time"

var limit = make(chan int, 3)

var works = []func(){
	func() {
		println("1")
		time.Sleep(1 * time.Second)
	},
	func() {
		println("2")
		time.Sleep(1 * time.Second)
	},
	func() {
		println("3")
		time.Sleep(1 * time.Second)
	},
	func() {
		println("4")
		time.Sleep(1 * time.Second)
	},
	func() {
		println("5")
		time.Sleep(1 * time.Second)
	},
}

func ChanEntry() {
	for _, w := range works {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select {}
}
