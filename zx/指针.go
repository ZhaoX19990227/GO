package main

func main() {
	a := 10
	println("a的地址为：", &a)
	var p *int
	p = &a
	println("p:", *p)

	var p2 **int
	p2 = &p
	println("p的地址：", &p)
	println("p2的值相当于是p的地址；", p2)

	println("p2的地址就是自己：", &p2)
	println("*p2其实就是p2存储的地址的数值，其实就是a的地址", *p2)
	println("p2地址对应的值就是a的值：", **p2)
}
