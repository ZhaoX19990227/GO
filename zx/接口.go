package main

import "fmt"

func main() {
	m1 := Mouse{"小胖"}
	test(m1)
}

// 接口
type USB interface {
	start()
	end()
}

// 实现类
type Mouse struct {
	name string
}

// 实现类
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("mouse start")
}
func (m Mouse) end() {
	fmt.Println("mouse end")
}

func (f FlashDisk) start() {
	fmt.Println("FlashDisk start")
}
func (m FlashDisk) end() {
	fmt.Println("FlashDisk end")
}

func test(usb USB) {
	usb.start()
	usb.end()
}
