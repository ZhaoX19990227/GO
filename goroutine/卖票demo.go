package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 全局变量 100张票
var ticket = 10

var wg sync.WaitGroup
var matex sync.Mutex

func main() {

	/**
	4个goroutine模拟4个售票窗口
	*/
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(3 * time.Second)
}

func saleTickets(name string) {
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("售出：", ticket)
			ticket--
		} else {
			fmt.Println("已售尽")
			break
		}
	}

}
