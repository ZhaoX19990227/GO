package main

import "fmt"

func main() {
	//p1 := new(User)
	//p1.name = "小胖"
	//p1.age = 24
	//p1.address.id = 1
	//p1.address.info = "江苏南京"
	//
	//fmt.Println(p1)

	p2 := new(Student)
	book := Book{
		title: "你好",
	}
	p2.book = &book
	p2.id = 2

	fmt.Println(p2)      //&{2 0x14000104220}
	fmt.Println(p2.book) //&{你好}
	fmt.Println(book)    //{你好}
	fmt.Println("==============")
	p2.book.title = "呼啸山庄"
	fmt.Println(book)    //{呼啸山庄}   修改后book对象也被修改了
	fmt.Println(p2)      //&{2 0x14000104220}
	fmt.Println(p2.book) //&{呼啸山庄}
}

//	type Address struct {
//		id   int
//		info string
//	}
//
//	type User struct {
//		name    string
//		age     int
//		address Address
//	}
type Student struct {
	id   int
	book *Book
}

type Book struct {
	title string
}
