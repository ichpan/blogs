// 查找算法
package main

import (
	"fmt"
)

// 顺序查找
func order(arr [6]int, element int) (idx int) {
	for idx, val := range arr {
		if val == element {
			return idx
		}
	}
	return -1
}

// 折半查找
func binary(arr *[6]int, left, right, val int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	if (*arr)[middle] > val {
		return binary(arr, left, middle-1, val)
	} else if (*arr)[middle] < val {
		return binary(arr, middle+1, right, val)
	} else {
		return middle
	}
}

func main() {
	arr := [...]int{1, 2, 3, 5, 8, 10}
	//idx := order(arr, 5)
	idx := binary(&arr, 0, len(arr)-1, 8)
	fmt.Println(idx)
}
