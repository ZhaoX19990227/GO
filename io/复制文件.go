package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	srcFile := "/Users/zhaoxiang/Desktop/zhaoxiang.txt"
	destFile := "/Users/zhaoxiang/Desktop/xiaopang.txt"
	//
	//total, err := copyFile(srcFile, destFile)
	//fmt.Println(total, err)

	total, err := copy3(srcFile, destFile)
	fmt.Println(total, err)
}

//func copyFile(srcFile, destFile string) (int, error) {
//	file1, err1 := os.OpenFile(srcFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err1 != nil {
//		fmt.Println("打开文件失败", err1)
//		return 0, err1
//	}
//	fmt.Println("file:", file1)
//	file2, err2 := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err2 != nil {
//		fmt.Println("打开文件失败：", err2)
//		return 0, err2
//	}
//	fmt.Println("file2:", file2)
//	defer file1.Close()
//	defer file2.Close()
//
//	bs := make([]byte, 1024, 1024)
//	total := 0
//	for {
//		n, err := file1.Read(bs)
//		if err == io.EOF || n == 0 {
//			fmt.Println("拷贝完毕")
//			break
//		} else if err != nil {
//			fmt.Println("出错了")
//			return total, err
//		}
//		total += n
//		//将【本次】读取的数据写出去 n个数据
//		file2.Write(bs[:n])
//	}
//	return total, nil
//}

//func copy2(srcFile, destFile string) (int64, error) {
//	file1, err1 := os.Open(srcFile)
//	if err1 != nil {
//		return 0, err1
//	}
//	file2, err2 := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
//	if err2 != nil {
//		return 0, err2
//	}
//	defer file1.Close()
//	defer file2.Close()
//	//参数反过来写 写是第一个参数 读是第二个参数
//	return io.Copy(file2, file1)
//}

func copy3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

//test
//123321
