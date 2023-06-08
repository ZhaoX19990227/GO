package main

import (
	"fmt"
	"os"
)

func main() {

	////如果文件不存在会创建，如果已经存在会将文件内容清空
	//file, err := os.Create("/Users/zhaoxiang/Desktop/zhaoxiang.txt")
	//if err != nil {
	//	fmt.Println("创建文件失败：", err)
	//	return
	//}
	//fmt.Println("创建文件成功", file)

	file1, err1 := os.Open("/Users/zhaoxiang/Desktop/zhaoxiang.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("打开成功：", file1)

	file2, err2 := os.OpenFile("/Users/zhaoxiang/Desktop/zhaoxiang.txt", os.O_RDWR, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	//file2.Close()
	fmt.Println("打开成功：", file2)
}
