package main

func main() {
	arr := [4]int{1, 2, 3, 4}

	var address = &arr
	println("address地址：", &address)

	//获取数组的值
	println((*address)[0])
}
