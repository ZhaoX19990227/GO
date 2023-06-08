package main

import (
	"log"
	"os"
)

func main() {
	openFile()
}
func openFile() {
	_, error := os.Open("text.txt")
	if error != nil {
		log.Fatal(error)
	} else {
		log.Fatal("打开文件成功")
	}
}
