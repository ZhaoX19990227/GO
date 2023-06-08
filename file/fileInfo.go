package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
)

func main() {
	fileInfo, err := os.Stat("/Users/zhaoxiang/Desktop/GO.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(fileInfo)) //*os.fileStat
	fmt.Println(fileInfo.Name())          //GO.md
	fmt.Println(fileInfo.Size())          //12897
	fmt.Println(fileInfo.Mode())          //-rw-r--r--
	fmt.Println(fileInfo.ModTime())       //2023-06-08 19:20:33.005291583 +0800 CST
	fmt.Println(fileInfo.IsDir())         //false
	fmt.Println(fileInfo.Sys())           //&{16777233 33188 1 26771363 501 20 0 [0 0 0 0] {1686223345 670160734} {1686223233 5291583} {1686223343 814918986} {1685891821 75042486} 12897 32 4096 64 0 0 [0 0]}

	//获取父目录
	parentPath := path.Join("/Users/zhaoxiang/Desktop/GO.md", "..") ///Users/zhaoxiang/Desktop
	fmt.Println(parentPath)

	err1 := os.Mkdir("/Users/zhaoxiang/Desktop/a", os.ModePerm)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("创建目录成功")
	//多层文件夹
	os.MkdirAll("/Users/zhaoxiang/Desktop/a/b/c", os.ModePerm)
}
