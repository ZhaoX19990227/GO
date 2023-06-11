package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)

	fun1()
	fun2()

	fmt.Println("main 进入阻塞状态，等待waitGroup中的子goroutine执行完成")
	waitGroup.Wait() //当waitGroup为0时则会唤醒这里的阻塞状态
	fmt.Println("main解除阻塞")
}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun1:A", i)
	}
	waitGroup.Done() //源码：wg.Add(-1)
}

func fun2() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("fun2...")
	}

}
