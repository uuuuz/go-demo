package main

import (
	"demo_test/reflect_demo/demo"
	"fmt"
	"reflect"
)

func main() {

	data := demo.NewStruct001()

	v := reflect.ValueOf(data)
	fmt.Println(v)
	v = v.Elem()
	v0 := v.FieldByName("name1")
	fmt.Println(v0)
	v01 := v0.FieldByName("flag1")
	fmt.Println(v01)
	fmt.Println(v01.Bool())
}
