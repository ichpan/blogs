package main

import (
	"fmt"
	"io"
	"net"
)

func Process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			return
		}
		fmt.Print(string(buf[:n]))
	}

}

func Server() {

	listen, err := net.Listen("tcp", "0.0.0.0:5535")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listen net.Listener) {
		err = listen.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(listen)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// 打印对方Ip
		// fmt.Println(conn.RemoteAddr())

		go Process(conn)
	}
}

func main() {
	Server()
}
