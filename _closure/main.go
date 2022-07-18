package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// 当多次调用的时候 return后面的不会重新被初始化

// 闭包案例
// 简单理解: 在golang中闭包就是一个类python

func makeSuffix(suffix string) func(filename string) string {

	return func(filename string) string {
		// 如果字符串以指定字符结尾,则符合规则，直接返回 反之
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix
		}
		return filename
	}
}

func main() {
	//f := AddUpper()
	//fmt.Println(f(1)) // 11
	//fmt.Println(f(3)) // 14
	f := makeSuffix(".png")
	fmt.Println(f("这是一张图片"))
	fmt.Println(f("这是一张图片.png"))
}
