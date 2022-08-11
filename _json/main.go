package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Birthday string `json:"birthday"`
	Salary   int    `json:"salary"`
	Skill    string `json:"skill"`
}

// 序列化 struct、slice、map
func SerializerStruct() string {
	monster := Monster{"牛魔王", 200, "2018-01-01", 5000, "健身"}
	// 序列化结构体
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func SerializerMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "古力娜扎"
	a["age"] = 30
	a["address"] = "内蒙古"
	// 序列化结构体
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func SerializerSlice() {
	var slice []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "古力娜扎"
	m["address"] = "北京"

	slice = append(slice, m)

	// 序列化结构体
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

// 反序列化
func DeserializationStruct() {
	data := SerializerStruct()
	var monster Monster
	err := json.Unmarshal([]byte(data), &monster)
	if err != nil {
		return
	}
	fmt.Println(monster)
}

func main() {
	//SerializerStruct()
	//SerializerMap()
	//SerializerSlice()
	DeserializationStruct()
}
