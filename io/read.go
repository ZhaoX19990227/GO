package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/zhaoxiang/Desktop/zhaoxiang.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close()
	fmt.Println("打开成功", file)

	//			 	  长度为4，容量为4
	bs := make([]byte, 4, 4)

	//bytes, err := file.Read(bs)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	return
	//}
	//fmt.Println(bytes)
	//fmt.Println(bs)

	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("文件读到末尾，结束")
			return
		}
		fmt.Println(string(bs[:n]))
	}
}
