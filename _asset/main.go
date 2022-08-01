package main

import "fmt"

type Point struct {
	x, y int
}

func test() {
	var a interface{}
	point := Point{1, 2}
	a = point

	var b Point
	// 类型断言,表示判断a是否指向Point类型的变量,如果是就转换成point类型并赋值给b变量,否则报错
	b = a.(Point)
	fmt.Println(b)

	var a1 interface{}
	var b1 int = 10
	a1 = b1
	c1 := a1.(int)
	fmt.Println(c1)
}

func AssetCheck() {
	var a interface{}
	var b int = 10
	a = b
	if c, success := a.(int); success {
		fmt.Println(c)
	} else {
		fmt.Println("转换失败！")
	}
}

// 判断传入参数的类型
func AssetType(items ...interface{}) {
	for _, val := range items {
		switch val.(type) {
		case int32:
			fmt.Println("1")
		default:
			fmt.Println("2")
		}
	}
}

func main() {
	// test()
	// AssetCheck()
	x := 32
	AssetType(x)
}
