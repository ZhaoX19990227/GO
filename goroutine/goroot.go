package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//fmt.Println(runtime.GOROOT()) ///Users/zhaoxiang/sdk/go1.19
	//
	//fmt.Println(runtime.GOOS) //darwin
	//
	//fmt.Println(runtime.NumCPU()) //8核

	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println(i)
	//	}
	//}()
	//for i := 0; i < 4; i++ {
	//	runtime.Gosched()
	//	fmt.Println("hello:", i)
	//}

	//创建goroutine
	go func() {
		fmt.Println("goroutine开始...")
		fun()
		fmt.Println("goroutine结束...")
	}()
	time.Sleep(3 * time.Second)
}
func fun() {
	//defer还是会执行
	defer fmt.Println("defer...")
	//使用return fun()不会打印 但是defer还是会执行
	//return
	//使用goexit()是对fun方法进行的goexit所以之后的方法都不会执行，但是defer还是会执行的 goroutine开始...  defer...
	runtime.Goexit()
	fmt.Println("fun()")
}
