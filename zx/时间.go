package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	format := t1.Format("2006-01-02 15:04:05")
	fmt.Println(format)
}
