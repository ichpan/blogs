package main

import "fmt"

// 声明接口

// 声明一个接口 模拟现实中的usb接口 实现接口插拔方法
type Usb interface {
	insert()
	remove()
}

// 声明一个结构体 phone
type Phone struct{}

// 插入方法
func (p *Phone) insert() {
	fmt.Println("设备已插入")
}

// 设备断开方法
func (p *Phone) remove() {
	fmt.Println("设备已移除")
}

type Computer struct{}

func (computer Computer) Working(usb Usb) {
	usb.insert()
	usb.remove()

}

// 基本数据类型实现接口
type integer int

func (i *integer) Speak() {
	fmt.Printf("Speak methods!\n")
}

type TestInterface interface {
	Speak()
}

func main() {
	//	phone := Phone{}
	//	computer := Computer{}
	//	computer.Working(&phone)

	var i integer = 10
	var b TestInterface = &i
	b.Speak()
}
