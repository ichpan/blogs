package main

import (
	"encoding/json"
	"fmt"
)

// 猫的结构体
type Cat struct {
	Name, Color string
	Age         int
}

func test() {
	var cat Cat
	cat.Name = "Tom"
	cat.Color = "Black"
	cat.Age = 2
	fmt.Println(cat)
}

type person struct {
	Name string
	Age  int
}

// 结构体定义
func Create() {
	// 直接定义
	var p0 person
	fmt.Println(p0)

	// 使用{}
	p1 := person{}
	fmt.Println(p1.Name)

	// 使用指针
	p2 := new(person)
	p2.Name = "ichpan"
	fmt.Println(*p2)

	// 使用{}
	p3 := &person{"123", 1}
	p3.Age = 10
	fmt.Println(p3.Age)
}

// 结构体相互转换
func transform() {
	type A struct {
		Num int
	}
	type B struct {
		Num int
	}
	var a A
	a.Num = 2
	var b B
	b.Num = 5
	a = A(b)
	fmt.Printf("%v", a.Num)
}

// 结构体tag
func Tag() {
	type Monster struct {
		Name  string `json:"name"`
		Skill string `json:"skill"`
		Age   int    `json:"age"`
	}

	obj := Monster{"熏悟空", "飞", 100}
	jsonStr, _ := json.Marshal(obj)
	fmt.Printf("%T", string(jsonStr))
}

func main() {
	//test()
	//Create()
	//transform()
	Tag()
}
