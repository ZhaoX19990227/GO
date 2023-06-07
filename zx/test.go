package zx

import "reflect"

func main() {
	name := "小胖"
	print(reflect.TypeOf(name).Name())
}
