package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {

	conn, err := net.Dial("tcp", "localhost:5535")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// stdout代表标准输入
	reader := bufio.NewReader(os.Stdout)
	for {
		line, _ := reader.ReadString('\n')
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("退出成功！")
			break
		}

		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func main() {
	Client()
}
