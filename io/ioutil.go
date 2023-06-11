package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileName := "/Users/zhaoxiang/Desktop/zhaoxiang.txt"
	//bytes, err := ioutil.ReadFile(fileName)
	//fmt.Println(err)
	//fmt.Println(string(bytes))
	fileName := "/Users/zhaoxiang/Desktop/zhaoxiang.txt"
	s1 := "春暖花开"

	err := ioutil.WriteFile(fileName, []byte(s1), os.ModePerm)
	fmt.Println(err)
}
