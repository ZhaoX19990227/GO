package main

import "fmt"

func main() {

	var channel chan bool

	channel = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i, "子goroutine...")
		}
		//往通道中写数据
		channel <- true
		fmt.Println("结束")
	}()
	//如果通道中没有数据则会进入阻塞
	data := <-channel
	fmt.Println(data)
}
