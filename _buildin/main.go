// 内建函数
package main

import "fmt"

func BuildInNew() {
	num := new(int)

	*num = 100
	fmt.Println(*num)
}

func main() {
	BuildInNew()
}
