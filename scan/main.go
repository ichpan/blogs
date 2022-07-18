package main

import (
	"fmt"
)

func main() {
	// 1、使用ScanLn
	var name string
	//fmt.Println("输入你的名字:")
	//fmt.Scanln(&name)
	//fmt.Println(name)

	// 2、使用Scanf 按照指定的格式
	// 以空格分开 可以格式化匹配
	num, err := fmt.Scan(&name)
	fmt.Println(num, err)
	fmt.Println(name)
}
