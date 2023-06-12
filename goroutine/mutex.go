package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//// 全局变量 10张票
//var ticket = 10
//
//// 创建锁头
//var mutex sync.Mutex
//var wg sync.WaitGroup
//
//func main() {
//
//	/**
//	4个goroutine模拟4个售票窗口
//	*/
//	wg.Add(4)
//	go saleTickets("售票口1")
//	go saleTickets("售票口2")
//	go saleTickets("售票口3")
//	go saleTickets("售票口4")
//	wg.Wait()
//
//	//time.Sleep(5 * time.Second)
//}
//
//func saleTickets(name string) {
//	//rand.Seed(time.Now().UnixNano())
//	defer wg.Done()
//	for {
//		mutex.Lock()
//		if ticket > 0 {
//			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
//			fmt.Println("售出：", ticket)
//			ticket--
//		} else {
//			mutex.Unlock()
//			fmt.Println("已售尽")
//			break
//		}
//		mutex.Unlock()
//	}
//}
