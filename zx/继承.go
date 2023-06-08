package main

import "fmt"

func main() {

	p1 := Person{
		name: "小胖",
		age:  22,
	}
	fmt.Println(p1)
	p2 := Male{
		school: "北京大学",
		Person: p1,
	}
	fmt.Println(p2)

	fmt.Println(p2.name)
}

type Person struct {
	name string
	age  int
}

type Male struct {
	school string
	Person // 匿名结构体字段 相当于继承了Person类
}
