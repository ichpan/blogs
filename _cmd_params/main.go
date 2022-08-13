package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run calc_test.go 参数1 参数2
	//params := os.Args
	//fmt.Printf("%v\n", params[1])

	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "root", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "IP")
	flag.IntVar(&port, "p", 3306, "端口")

	flag.Parse()
	fmt.Println(user, pwd, host, port)
}
