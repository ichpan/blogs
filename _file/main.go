package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 文件打开关闭
func OpenFile() {
	file, _ := os.Open("/Users/ichpan/blogs/_file/test.txt")

	// 输出文件 point类型
	fmt.Println(file)

	err := file.Close()
	if err != nil {
		fmt.Println("文件关闭失败！")
	}
}

// 带缓冲区的读文件
func Buffer(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败！")
		}
	}(file)

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	fmt.Println("文件读取结束！")
	return nil
}

// 一次性读取文件 文件较小的情况下
func Single(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func main() {
	//OpenFile()
	//err := Buffer("_file/test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	Single("_file/test.txt")
}
