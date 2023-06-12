package main

import "fmt"

func main() {
	var channel chan int
	fmt.Println("%T,%v\n", channel, channel)

	if channel == nil {
		fmt.Println("channel为空，先创建通道")
		channel = make(chan int)
		fmt.Println(channel)
	}
	test(channel)
}

func test(channel chan int) {
	fmt.Println("%T,%v\n", channel, channel)
}
