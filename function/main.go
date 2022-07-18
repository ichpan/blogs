// 函数学习
// 为了减小代码冗余
package main

import "fmt"

// 函数语法
/*
func 函数名 (形参列表) (返回值列表) {
	执行语句

	return 返回值列表
}

可变参数 args... int
args是slice 通过索引可以取值

init函数可以初始化一些参数
*/

// Fibonacci 递归
// 在一个函数中直接或间接的调用自己的操作
func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// GetResult 求出f(n) = 2 * f(n-1) + 1
func GetResult(n int) int {
	if n == 1 {
		return 3
	}
	return 2*GetResult(n-1) + 1
}

/*
猴子吃桃子问题
有一堆桃子,猴子第一天吃了其中的一半,并再多吃了一个,当到第十天时只有一只桃子了。问最初有多少桃子
*/

func MonkeyEatPeach(n int) int {
	if n == 10 {
		return 1
	}
	return (MonkeyEatPeach(n+1) + 1) * 2
}

func main() {
	num := Fibonacci(6)
	fmt.Println(num)

	//num2 := GetResult(5)
	//fmt.Println("num:", num2)

	//fmt.Println(MonkeyEatPeach(1))
}
