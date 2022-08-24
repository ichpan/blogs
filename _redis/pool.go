package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 32,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}}
}

func CloseConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatalf("关闭链接失败！")
	}
}

func main() {
	conn := pool.Get()
	defer CloseConn(conn)

	_, err := conn.Do("hmset", "student", "name", "小王", "age", 18)
	if err != nil {
		log.Fatalf("写入失败！")
	}
}
