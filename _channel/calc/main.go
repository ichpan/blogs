package main

import "fmt"

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, PrimeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			PrimeChan <- num
		}
	}

	exitChan <- true
}

func main() {
	intChan := make(chan int, 2000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
