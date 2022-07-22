package main

import "fmt"

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

func main() {
	//test()
	Create()
}
