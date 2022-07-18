package main

import (
	"fmt"
	"time"
)

// 获取当前时间
func GetCurrentTime() {
	now := time.Now()
	fmt.Printf("now的值是：%v 类型是: %T \n", now, now)

	fmt.Printf("%v \n", int(now.Month()))
}

// 格式化日期时间
func FormatDate() {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(date)

	CurrentTime := now.Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间是:%v \n", CurrentTime)

	fmt.Println(now.Unix())
}

func main() {
	GetCurrentTime()
	FormatDate()
}
