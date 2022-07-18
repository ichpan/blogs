package main

import (
	"blogs/_format"
	"fmt"
)

func main() {
	err := _format.FmtString()
	fmt.Println(err)
}
