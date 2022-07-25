package main

import "fmt"

// 表示结构体A有一个test方法
type Person struct {
	age int
}

// (p Person)体现绑定
func (a Person) test() {
	a.age = 65535
	fmt.Println(a.age)
}

// 绑定一个speak方法
func (a Person) speak() {
	fmt.Println("age is ", a.age)
}

// 添加一个计算的方法
func (a Person) calc(n int) (res int) {
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

// 计算圆的面积
type Circle struct {
	radius float64
}

// 为Circle编写一个方法,返回面积
func (c *Circle) area(radius float64) float64 {
	c.radius = radius
	return 3.14 * c.radius * c.radius
}

func main() {
	// 调用
	var p Person
	p.test()
	p.speak()
	fmt.Println("num is ", p.calc(3))

	var c Circle
	area := c.area(1)
	fmt.Println(area)
}
