package main

import (
	"errors"
	"fmt"
	"net"
	"reflect"
)

func main() {
	////方法一
	//err := errors.New("自定义错误")
	//fmt.Println(err)
	//
	////方法二
	//err1 := fmt.Errorf("自定义错误，错误码：%d", 100)
	//fmt.Println(err1)
	println(check(10))

	_, err := net.LookupHost("www.baidu.com")
	if ins, ok := err.(*net.DNSError); ok {
		if ins.IsTimeout {
			fmt.Println("超时")
		} else if ins.IsNotFound {
			fmt.Println("未找到")
		} else {
			fmt.Println("其他错误")
		}
	}

}

func check(age int) error {
	if age < 0 || reflect.TypeOf(age) == nil {
		return errors.New("年龄有误")
	} else {
		return nil
	}
}
