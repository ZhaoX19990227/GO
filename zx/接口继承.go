package main

import "fmt"

func main() {

	var cat Cats = Cats{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("---------------")
	var a1 A1 = cat
	a1.test1()

	var b1 B1 = cat
	b1.test2()

	var c1 C1 = cat
	c1.test1()
	c1.test2()
	c1.test3()

}

type A1 interface {
	test1()
}

type B1 interface {
	test2()
}

type C1 interface {
	A1
	B1
	test3()
}
type Cats struct {
}

func (c Cats) test1() {
	fmt.Println("test1")
}
func (c Cats) test2() {
	fmt.Println("test2")
}

func (c Cats) test3() {
	fmt.Println("test3")
}
