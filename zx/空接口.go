package main

import "fmt"

func main() {
	var a1 = Cat{color: "123"}
	fmt.Println(a1)
	testColor(a1)
	map1 := make(map[string]interface{})
	map1["name"] = "小花"
	map1["age"] = 30
	map1["Cat"] = Cat{
		color: "红色",
	}
	fmt.Println(map1)

	//定义切片 空接口类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, "1")
	slice1 = append(slice1, "2")
	slice1 = append(slice1, "3")
	//循环切片内容
	testSlice(slice1)
}

type A interface {
}

type Cat struct {
	color string
}

func testColor(a A) {
	fmt.Println(a)
}

func testSlice(slice1 []interface{}) {
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
}
