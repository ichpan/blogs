package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 数组的声明
func statement() {
	var names [3]string
	fmt.Println(cap(names))

	ages := [...]int{}
	fmt.Printf("%v", ages)
}

// 数组的便利
func cycle() {
	nums := [...]int{1, 3, 5}
	// 普通便利
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
	}
	fmt.Println()

	// range使用
	for idx, val := range nums {
		fmt.Println(idx, val)
	}
}

// 数组函数
func Test(arr *[3]int) bool {
	arr[0] = 0
	if arr[0] == 0 {
		return true
	}
	return false
}

// 案例
// 创建一个byte类型的26个元素的数组,分别放置A-Z。使用for♻️打印出来
func Demo1() {
	MyCharsArray := [26]byte{}
	for i := 0; i < len(MyCharsArray); i++ {
		MyCharsArray[i] = 'A' + byte(i)
		fmt.Printf("%c", MyCharsArray[i])
	}
}

// 求出一个数组的最大值
func Demo2() float64 {
	intArray := [...]float64{1, 4, 2, 7, 9, 4}
	max := intArray[0]
	for _, val := range intArray {
		max = math.Max(max, val)
	}
	return max
}

// 求出一个数组的和和平均值
func Demo3() {
	intArray := [...]int{6, 4, 2, 7, 9}
	sum := 0
	for _, val := range intArray {
		sum += val
	}
	avg := float64(sum) / float64(len(intArray))
	fmt.Printf("%v, %v", sum, avg)
}

// 随机生成5个数,并反转打印
func Demo4() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	nums := [5]int{}
	for i := 0; i < 5; i++ {
		nums[4-i] = rand.Intn(100)
	}
	fmt.Println(nums)
}

func main() {
	//statement()
	//cycle()

	// 修改数组的值 数组是值类型,想要修改数组的值,需要传递指针
	//arr := [3]int{1, 2, 3}
	//Test(&arr)
	//fmt.Println(arr)

	//Demo1()
	//fmt.Println(Demo2())
	//Demo3()
	Demo4()
}
