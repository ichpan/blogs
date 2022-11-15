package main

import "fmt"

func calc(c chan bool) int {
	total := 0
	for i := 0; i < 10000; i++ {
		total += i
	}
	c <- true
	fmt.Printf("%v", total)
	return total
}

func main() {
	c := make(chan bool)
	go calc(c)

	// 同步等待
	<-c
}
