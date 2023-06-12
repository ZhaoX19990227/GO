package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个定时器 3秒后触发
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("%T\n", timer)

	fmt.Println(time.Now())
	//从通道中获取值
	ch := timer.C
	fmt.Println(<-ch)
}
