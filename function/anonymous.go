// 匿名函数
// 匿名函数就是没有名字的函数

/*
func (形参) 返回值 {
	return val
}(实参)

也可以为函数赋值给变量 通过变量调用

全局匿名函数：
var (
	Anony = func (n1 int,n2 int) int {
		return n1 * n2
	}
)
*/
package main

import "fmt"

// Anon 全局匿名函数计算两束的积
var (
	Anon = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	val := Anon(1, 3)
	fmt.Println(val)
}
