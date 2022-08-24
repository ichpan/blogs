package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func ConnRedis() redis.Conn {
	// 建立链接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalf("连接redis失败！")
	}
	return conn
}

// 关闭链接
func CloseRedis(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("关闭链接失败！")
	}
}

func main() {
	r := ConnRedis()
	defer CloseRedis(r)
	_, err := r.Do("set", "name", "去洗碗啊！")
	if err != nil {
		log.Fatalf("%v", err)
	}

	res, _ := redis.String(r.Do("get", "name"))
	fmt.Println(res)
}
