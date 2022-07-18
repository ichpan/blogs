// 基本数据类型转换

package main

import (
	"fmt"
	"strconv"
)

// 转string
func test() {
	// 数据转换的时候需要注意防止溢出
	// Sprintf采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。

	var num1 int = 99
	//var num2 float64 = 99
	//var a bool = true
	//var b byte = 'A'
	var s string = "ichpan"

	// 转str int %d十进制  %f 浮点型 %t 布尔 %c char    Itoa方法将int转string
	//s = fmt.Sprintf("%d", num1)
	//fmt.Println(s)

	// 使用strconv 函数
	s = strconv.FormatInt(int64(num1), 10)
	fmt.Print(s)

}

// 基本转string
// str -> int 确保数据类型合适 转换失败返回0
func test1() {
	s := "123"
	// Atoi 将string转int
	//num, _ := strconv.ParseInt(s, 10, 0)
	num, _ := strconv.Atoi(s)
	fmt.Print(num)
}

func main() {
	test1()
}
