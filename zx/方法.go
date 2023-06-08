package main

import "fmt"

func main() {
	//通过对象调用
	w1 := Worker{
		name: "小胖",
		age:  24,
	}
	w1.work()
	//通过指针调用
	w2 := &Worker{
		name: "小猪佩奇",
		age:  25,
	}
	w2.work()
	w2.rest()
}

type Worker struct {
	name string
	age  int
}

func (work Worker) work() {
	fmt.Println(work.name, "在工作")
}

func (p *Worker) rest() {
	fmt.Println(p.name, "在休息")
}
