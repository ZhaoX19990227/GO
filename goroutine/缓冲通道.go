package main

import "fmt"

func main() {
	channel := make(chan int, 10)
	fmt.Println(len(channel), cap(channel))
}
