// 排序算法
package main

// 冒泡排序
func bubble(arr *[5]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

// 递归 冒泡排序
func ReBubble(arr *[5]int) {
	flag := false
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			flag = true
		}
	}
	if flag {
		ReBubble(arr)
	}
}

//func main() {
//	arr := [5]int{1, 6, 9, 2, 5}
//	ReBubble(&arr)
//	fmt.Println(arr)
//}
