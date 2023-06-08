package main

import (
	"fmt"
	"os"
)

func main() {
	//只有使用append才会追加写入 不加的话是从头开始写入
	file, err := os.OpenFile("/Users/zhaoxiang/Desktop/zhaoxiang.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	fmt.Println(file)
	//defer file.Close()
	bs := []byte{65, 98, 93, 78, 68}
	n, err1 := file.Write(bs)
	if err1 != nil {
		fmt.Println("写入文件失败：", err1)
		return
	}
	fmt.Println(n)
}
