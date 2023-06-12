package main

import "fmt"

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan<- int) //只能写入
	//ch3 := make(<-chan int) //只能读取

	ch1 <- 100
	var ch1data = <-ch1
	fmt.Println(ch1data)

	//ch2 <- 200
	//fmt.Println(ch2)

	//print(ch3)
}
