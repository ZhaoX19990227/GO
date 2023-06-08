package main

import "fmt"

func main() {
	//方法一
	var p1 person
	p1.name = "赵翔"
	p1.sex = "男"
	p1.age = 24
	p1.address = "江苏"

	////方法二
	//p2 := person{}
	//p2.name = "小胖"
	//
	////方法三
	//p3 := person{"小胖", 22, "男", "江苏南京"}

	//结构体指针
	//var pp1 *person
	//pp1 = &p1
	//fmt.Println(pp1)
	////值
	//fmt.Println(*pp1)
	//
	////会将引用的p1也修改了 说明结构体指针是引用类型
	//(*pp1).name = "余晓桅"
	////可以简写 pp1.name
	//fmt.Println(*pp1)
	//fmt.Println(p1)

	p2 := new(person)
	fmt.Println(p2)
	(*p2).name = "Tom"
	//简写
	p2.name = "Lucy"
}

type person struct {
	name    string
	age     int
	sex     string
	address string
}
