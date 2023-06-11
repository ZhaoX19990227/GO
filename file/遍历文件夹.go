package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "/Users/zhaoxiang/IdeaProjects/myleetcode"
	listFiles(dirname, 0)
}

func listFiles(dirname string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range fileInfos {
		filename := dirname + "/" + info.Name()
		fmt.Printf("%s%s\n", s, filename)
		if info.IsDir() {
			listFiles(filename, level+1)
		}
	}
}
