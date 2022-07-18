// 数据类型
// byte计算机中基本存储单元 bit计算机中最小存储单位

// 基本数据类型
// 数值  整数、浮点   rune int32的别名    complex32
// 整型 默认为int
// 字符型 byte(保存单个字母) 0-255
// 布尔型 true / false
// 字符串 string

// 派生/复杂数据类型
// 指针
// 数组 array
// 结构体 struct
// 管道 channel
// 函数 func
// 切片 slice
// 接口 interface
// map

package main

import (
	"fmt"
	"unsafe"
)

// uint为无符号数

func IntType() {
	// 2^n-1
	// int8 占用1个字节
	// int16 占用2个字节
	// int32 占用4个字节
	// int64 占用8个字节
	var intNum int = 12
	fmt.Println(intNum)

	//年龄
	var age byte = 21
	fmt.Println(age)

	// 格式化输出 输出类型
	fmt.Printf("%T \n", intNum)

	// 输出占用字节大小unsafe.Sizeof()
	fmt.Printf("%d", unsafe.Sizeof(age))
}

// 浮点类型
func typeFloat() {
	// 存放一个价格
	// 浮点数可能会有精度损失  64比32更准确  默认为float64
	price := -8.21345
	fmt.Printf("%T", price)
}

// 字符类型 使用utf-8编码
// utf-8 英文是一个字节 汉字是3个字节 unicode就是
func typeChar() {
	// 如果要存储一个字母，推荐使用byte 范围0-255
	// 0-9 A-Z a-z
	var val int = 'A'
	// 该值对应的unicode码值
	fmt.Println(val)
}

// 布尔类型
func typeBool() {
	// 占用字符为1个 默认值默认值为false
	var val = false
	fmt.Println(val, unsafe.Sizeof(val))
}

// 字符串
func typeString() {
	// 使用`` 原样输出
	// 字符串拼接的时候换行的时候 + 保留在上一行的位置
}

func main() {
	//IntType()
	//typeFloat()
	//typeChar()
	//typeBool()
	//typeString()
}
