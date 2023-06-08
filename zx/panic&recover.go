package main

import "fmt"

func main() {
	defer func() {
		//if msg := recover(); msg != nil {
		//	fmt.Println(msg, "成功恢复")
		//}

	}()
	fun()
}

func fun() {

	for i := 0; i < 10; i++ {
		defer fmt.Println("123")
		if i == 5 {
			//panic("i==5恐慌，后续代码不会执行，但是之前defer的会执行\n")
		}
	}
	defer println(456)
}
