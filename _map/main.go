package main

import (
	"fmt"
	"sort"
)

// map声明
func statement() {
	dict := make(map[string]string, 2)
	dict["name"] = "ichpan"

	heroes := map[int]string{
		1: "宋江",
		2: "张飞",
	}
	heroes[3] = "李四"
	fmt.Println(heroes)
}

// map的查找
func find() {
	heroes := map[int]map[string]string{
		1: {"name": "宋江"},
		2: {"name": "张飞"},
	}
	if heroes[3] != nil {
		fmt.Println("exist!")
	}
	val, exist := heroes[6]
	if exist {
		fmt.Println(val)
	}
}

// map便利
func cycle() {
	heroes := map[int]string{
		1: "宋江",
		2: "张飞",
	}
	for key, val := range heroes {
		fmt.Println(key, val)
	}
}

// map切片
func mapSlice() []map[string]string {
	msg := make([]map[string]string, 0)
	student := map[string]string{
		"name": "张三",
		"age":  "18",
	}
	msg = append(msg, student)
	return msg
}

// map排序
func mapSorted() {
	map1 := map[int]string{
		1: "张三",
		4: "李四",
		3: "王五",
	}
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Println(map1[key])
	}
}

func main() {
	//statement()
	//find()
	//cycle()
	//fmt.Println(mapSlice())
	//mapSorted()
}
