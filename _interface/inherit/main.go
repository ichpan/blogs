// 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋值给接口
package main

import "fmt"

type A interface {
	test01()
}

type B interface {
	test02()
}

type C interface {
	A
	B
	test03()
}

// 学生的结构体
type student struct {
}

func (stu *student) test01() {}
func (stu *student) test02() {}
func (stu *student) test03() {}

func main() {
	var stu student
	var c C = &stu
	c.test01()

	// 空接口没有任何方法，可以为它赋值任何类型
	num := 10
	var test interface{} = num
	fmt.Println(test)
}
