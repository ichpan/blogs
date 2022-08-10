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

// 文件读写
func ReadWrite() {
	file, err := os.OpenFile("_file/test2014.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	// 读取文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	// 写入内容
	str := "hello world123!\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

// 判断文件是否存在
func FileIsExist() (bool, error) {
	_, err := os.Stat("_file/test1.txt")
	if err != nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}

// 文件拷贝
func CopyFile(dst, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

// 记录英文、数字、空格、其他字符的个数
type CharCount struct {
	EnCount, NumCount, SpaceCount, OtherCount int
}

// 统计文件中字符的个数
func CountChar(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, val := range str {
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough
			case val >= 'a' && val <= 'z':
				count.EnCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			case val >= 0 || val <= 9:
				count.SpaceCount++
			case val == ' ' || val == '\t':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	//OpenFile()
	//err := Buffer("_file/test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//Single("_file/test.txt")
	//ReadWrite()
	//exist, err := FileIsExist()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(exist)

	filepath := "_file/test.txt"
	CountChar(filepath)
}
