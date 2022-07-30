// 一个自定义类型可以实现多个接口
package main

import "fmt"

type A interface {
	Speak()
}

type B interface {
	Working()
}

// 空接口people
type People struct{}

func (peo *People) Speak() {
	fmt.Println("Speak！")
}

func (peo *People) Working() {
	fmt.Println("Working")
}

func main() {
	var p People
	var a A = &p
	var b B = &p
	a.Speak()
	b.Working()
}
