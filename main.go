package main

func main() {
	//result := fun.Hello()
	//println(result)
	//
	//var (
	//	name = "冷少"
	//)
	//println(name + "是志鹏的老公")
	//_, i := fun.GetNameAndAge()
	//println(i)
	const PI = 3.14
	const PI2 float64 = 3.1415
	const (
		name = "你好"
		age  = 12
	)
	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	print(a1, "", a2, "", a3)

}
