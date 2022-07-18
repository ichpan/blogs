// 变量使用

// var 变量名
//
// 可以一次性定义多个变量 var ()
// 可以使用 :代替 var   := 自动匹配类型 定义赋值

// int float 默认都为0 string默认为空字符串

package main

import "fmt"

func variable() string {
	var name string = "ichpan"
	name = "123"
	return name
}

// + 的使用
func add() string {
	var (
		x = "1"
		y = "2"
	)

	return x + y
}

func main() {
	val := variable()
	fmt.Print(val)
}
