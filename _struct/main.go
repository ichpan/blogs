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

func main() {
	test()
}
