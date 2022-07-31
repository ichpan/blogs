// 接口和继承的对比
package main

import "fmt"

// 猴子结构体
type Monkey struct {
	Name string
}

// 小猴子结构体
type LittleMonkey struct {
	Monkey
}

// 爬树
func (monkey *Monkey) climb() {
	fmt.Printf("%v会爬树！\n", monkey.Name)
}

type BirdAble interface {
	fly()
}

// 飞
func (little LittleMonkey) fly() {
	fmt.Println("会飞呢~")
}

func main() {
	m := LittleMonkey{
		Monkey{
			"熏悟空",
		},
	}
	m.climb()
	m.fly()
}
