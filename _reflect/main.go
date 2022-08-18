package main

import (
	"fmt"
	"reflect"
)

// 对基本数据类型进行反射的基本操作
func ReflectTest(val interface{}) {
	rTyp := reflect.TypeOf(val)
	fmt.Println(rTyp)

	// 真实的类型不是int 是一个reflect.Value类型
	rVal := reflect.ValueOf(val)
	num := rVal.Int() + 2
	fmt.Println(num)

	// 继续转成interface{}
	iv := rVal.Interface()
	num2 := iv.(int)
	fmt.Println(num2)
}

// 对struct反射的基本操作
type Student struct {
	Name string
	Age  int
}

func ReflectStruct(val interface{}) {

	rTyp := reflect.TypeOf(val)
	fmt.Println(rTyp)

	rVal := reflect.ValueOf(val)
	rv := rVal.Interface()
	stu, ok := rv.(Student)
	if ok {
		fmt.Println(stu.Name, stu.Age)
	}
}

func main() {
	num := 100
	ReflectTest(num)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	ReflectStruct(stu)
}
