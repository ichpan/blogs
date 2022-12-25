package main

import "fmt"

// 计算1-100的和
func calc(c chan bool) {
	total := 0
	for i := 0; i <= 100; i++ {
		total += i
	}
	fmt.Printf("%v", total)
	c <- true
}

func main() {
	c := make(chan bool)
	go calc(c)

	// 同步等待
	<-c
}
