package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)
	go send(channel)

	for v := range channel {
		fmt.Println(v)
	}

}

func send(channel chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		channel <- i
	}
	close(channel)
}
