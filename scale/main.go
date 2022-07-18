// 进制
// 二进制 0,1
// 八进制 0,1,2,3,4,5,6,7,7
// 十进制 0,1,2,3,4,5,6,7,8,9
// 十六进制 0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F

package main

import "fmt"

// 转换 格式化输出的时候
// 0b开头 二进制
// 0开头 八进制
// 十进制 默认
// 0x 十六进制

func transform() {
	number := 0x11
	fmt.Printf("%b", number)
}

// TwoToTen 二进制转十进制
// 2^n = 1+N个0
// ！！！规则   从最低位开始(右边)，将每个位上的数提取出来，乘以2的(位数-1)次方，求和
func TwoToTen() {
	fmt.Println(-2 & 3)
}

func main() {
	//transform()
	TwoToTen()
}
