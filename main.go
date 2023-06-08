package main

func main() {
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
