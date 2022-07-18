package main

import "fmt"

func TestDefer() {
	defer fmt.Println("程序执行完成!")
	defer fmt.Println("程序执行完成1!")
	defer fmt.Println("程序执行完成2!")
	fmt.Println("Defer主函数")
}

func main() {
	TestDefer()
}
