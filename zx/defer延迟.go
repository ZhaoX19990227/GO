package main

import "fmt"

func main() {

	defer fmt.Println("123")
	fmt.Println("234")
	defer fmt.Println("345")
	fmt.Println("456")

}

func fun1(str string) {
	fmt.Println("str...")
}
