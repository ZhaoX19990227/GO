package main

import "fmt"

func main() {
	chan1 := make(chan int)
	go sendData(chan1)

	for {
		v, ok := <-chan1
		if !ok {
			fmt.Println("已经获取了所有数据", ok)
			break
		}
		fmt.Println(v)
	}
}

func sendData(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
