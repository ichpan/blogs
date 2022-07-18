package main

import "fmt"

func ErrorTest() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func Test() int {
	defer ErrorTest()
	num1 := 10
	num2 := 0
	res := num1 / num2
	return res
}

func main() {
	Test()
	fmt.Println("发生错误!")
}
