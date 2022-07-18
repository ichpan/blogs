// 分支控制 单分支/多分支
// if 条件表达式 {    条件为真的时候执行
//    执行代码块
// }

package main

import "fmt"

func CalcAge() {
	// 获取用户输入
	//var age int
	//fmt.Scanln(&age)

	if age := 1; age > 10 {
		fmt.Println("年龄大于十岁！")
	} else if age == 10 {
		fmt.Println("年龄等于10岁")
	} else {
		fmt.Printf("小于10岁")
	}

}

func SwitchCase() {
	var val byte
	fmt.Println("\n请输入一个字符")
	fmt.Scanf("%c", &val)
	switch val {
	case 'a':
		fmt.Println("a")
	case 'b':
		fmt.Println("b")
	case 'c':
		fmt.Println("c")
	case 'd':
		fmt.Println("d")
	default:
		fmt.Println("Error")
	}

}

// for循环

// for{}  是死循环
func ForTest() {

	for i := 0; i <= 10; i++ {
		fmt.Println("测试")
	}

	str := "string"
	for index, val := range str {
		fmt.Printf("%d, %c, \n", index, val)

	}
}

// while
func test() {
	var i int = 1
	for {
		if i >= 10 {
			break
		}
		i++
		fmt.Println(123)
	}
}

// do while i++ 在判断前面

// 尽量避免使用goto语句，防止流程混乱
// goto语句可以无条件的转移到程序中指定的行
// 使用的时候我们需要定义标签即可
func TestGoto() {

}

func main() {
	//CalcAge()
	//SwitchCase()
	ForTest()
}
