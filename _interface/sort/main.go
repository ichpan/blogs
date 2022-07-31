// 结构体切片排序(按照分数)
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name  string
	Grade int
	Score int
}

type student []Student

func (stu student) Len() int {
	return len(stu)
}

func (stu student) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}

func (stu student) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var s student
	for i := 0; i <= 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("Name%d", rand.Intn(100)),
			Grade: 3,
			Score: int(rand.Float64() * 100),
		}
		s = append(s, stu)
	}
	sort.Sort(s)
	fmt.Println(s)
}
