# Map

> map是Key-Value数据结构,又称为字段或者关联数组,类似集合。
> map是引用类型

**基本语法:**
~~~go
// 申明
make(map[string]string, num)
// 直接赋值
map[int]string{}
~~~
**CURD：**
~~~go
// 统计长度
len(map)
// 增加/修改(key不存在则是新增/存在就是修改)
map[key] = value
// 删除
// 内建函数delete按照指定的键将元素从映射中删除。
// 若m为nil或无此元素，delete不进行操作。
delete(map, key)
直接清空可以给原来的map分配一个没有元素的空间
// 查找 如果exist是true则存在，反之, 特殊类型可以用nil来判断
val, exist := map[key]

~~~
**便利**
~~~go
// golang的便利只能使用range便利
for key, val := range map {
    ...
}
~~~
**切片**
~~~go
// 和切片一样
make([]map[string]string, 0)
~~~
**排序**
~~~go
使用sort包中带的方法,利用切片
~~~
**使用注意**
- map的容量到达后会自动扩容,可以动态变化
- map的value经常使用到struct