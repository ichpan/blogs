## 数组

> 数组是值类型,可以存放多个同一类型的数据

**数组内存布局：**

| 0 | 1 | 2 | 3 |
|---|---|---|---|

**数组的定义:**

~~~go
// 只声明数组
var array [num]int
// 直接赋值
var array [num]int = [num]int {1, 2, 3}
// 类型推断
var array = [num]int {1, 2, 3}
// 不指定大小
var array = [...]int {1, 2, 3}
// 指定元素对应的下标
var array = [3]string{1:"i", 2:"c", 3:"h"}
~~~
**数组的便利**
~~~go
# 基本便利
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}

# for-range结构便利
for index, value := range array {
}
~~~
**数组使用注意事项**
- 数组是多个类型数据的结合,一个数组一旦被定义,长度是固定的,不能动态变化
- var array []int 这个是slice不是数组
- 数组的元素可以是任意类型的元素,不可以混用
- 数组定义好之后,如果没有赋值,默认为零值(int、float -> 0;bool -> false,string -> " ")
- 数组索引是从0开始的
- 数组索引必须在指定范围内使用,不然会产生panic
- 长度是数据类型的一部分,传参的时候需要考虑
