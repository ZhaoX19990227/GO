package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "/Users/zhaoxiang/Desktop/zhaoxiang.txt"
	file, err := os.Open(fileName)
	HandleError(err)

	//defer file.Close()

	b1 := bufio.NewReader(file)
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	HandleError(err)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))
}

func HandleError(error error) {
	if error != nil {
		fmt.Println(error)
		return
	}
}
