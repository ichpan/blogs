package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func statement() {
	mychan := make(chan interface{}, 3)
	mychan <- 5

	cat := Cat{"大白", 5}
	mychan <- cat

	<-mychan
	mycat := <-mychan

	// asset
	a := mycat.(Cat)
	fmt.Printf("%v", a.Name)

	// close
	close(mychan)
}

func Range() {
	mychan := make(chan interface{}, 3)
	mychan <- 5

	cat := Cat{"大白", 5}
	mychan <- cat

	// close
	close(mychan)

	for v := range mychan {
		fmt.Println(v)
	}
}

func main() {
	statement()
	Range()
}
