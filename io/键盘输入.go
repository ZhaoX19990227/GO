package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//readString, _ := reader.ReadString('\n')
	//fmt.Println(readString)

	fileName := "/Users/zhaoxiang/Desktop/abc.txt"
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	//defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10000; i++ {
		writer.WriteString(fmt.Sprintf("%d,Hello", i))
	}
	writer.Flush()
}
