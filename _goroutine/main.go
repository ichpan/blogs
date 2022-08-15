package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test01() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello World " + strconv.Itoa(i))
		time.Sleep(1)
	}
}

// 计算1-20的各个数的阶乘，将结果放入map中，最后显示出来，使用goroutine完成
// 可能会产生并发/并行安全问题

var (
	mapping = make(map[int]int, 10)
	// 全局互斥锁
	lock sync.Mutex
)

func CalcFactorial(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

func test(key int) {
	mapping[key] = CalcFactorial(key)
}

func main() {
	//CpuNum := runtime.NumCPU()
	//runtime.GOMAXPROCS(CpuNum)

	time.Sleep(time.Second)

	for i := 1; i <= 20; i++ {
		lock.Lock()
		go test(i)
		lock.Unlock()
	}
	fmt.Println(mapping)
}

//func main() {
//go test()
//
//for i := 0; i <= 10; i++ {
//	fmt.Println("Hello Golang " + strconv.Itoa(i))
//	time.Sleep(1)
//}

//}
