package main

import (
	"fmt"
	"strconv"
	"time"
)

func state() {
	// 声明管道为只写
	OnlyWriteChan := make(chan<- int, 2)
	fmt.Println(OnlyWriteChan)

	// 声明为只读
	OnlyReadChan := make(<-chan int, 2)
	fmt.Println(OnlyReadChan)
}

func SelectUse() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		stringChan <- "hello-" + strconv.Itoa(i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case v := <-stringChan:
			fmt.Println(v)
		default:
			fmt.Println("取不到了！")
			return
		}
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello")
	}
}

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error")
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	//state()
	//SelectUse()

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("main(%d)", i)
	}

}
