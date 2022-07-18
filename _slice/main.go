// slice的基本使用
package main

import "fmt"

// 创建切片
func CreateSlice() {
	// 定义
	var slice = make([]int, 5)
	fmt.Println(cap(slice))

	// 直接赋值定义切片
	var demo = []int{1, 2, 3}
	demo[2] = 3
	fmt.Printf("%v", demo)
}

// 切片便利
func SliceConv() {
	// 常规
	var slice = make([]int, 5)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for-range
	for _, val := range slice {
		fmt.Println(val)
	}
}

// 操作切片
func UseSlice() []int {
	sliceCase := []int{1, 3, 5, 7, 9}
	sliceCase = append(sliceCase, 8)
	return sliceCase
}

// 切片拷贝
func SliceCopy() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Println(slice2)
}

func test() {
	a := make([]int, 5)
	a = append(a, 1, 3, 9, 6, 7)
	slice := make([]int, 6)
	copy(slice, a)
	fmt.Println(slice)
}

// string和slice的区别
func stringSlice() (arr []rune) {
	str := "hello world 中国"
	arr = []rune(str)
	arr[0] = 'i'
	return arr
}

func fbn(n int) (arr []int) {
	arr = make([]int, n)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func main() {
	// CreateSlice()
	// SliceConv()
	//fmt.Println(UseSlice())
	//SliceCopy()
	//test()
	//bytes := stringSlice()
	//fmt.Println(string(bytes))
	fmt.Println(fbn(10))
}
