# 面向对象
> &emsp;&emsp;GoLang是基于**Struct**来实现面向对象特性的,具有封装、继承、多态的特性(继承是通过匿名字段实现的)

## Struct的内存布局
- struct是自定义的数据类型,代表一类事务
- struct是具体的,代表一个具体变量
- struct是值类型的
- 不同结构体变量的字段是独立的
- 在创建一个结构体之后,还没有赋值的情况下,对应零值
- 指针、slice、map的零值都为nil,使用的时候需要分配内存
- .的优先级比*高 所以 *name.val是错误的

![](/Users/ichpan/blogs/asset/结构体内存图.png)

## 结构体申明
~~~go
// 基本语法
// 字段是结构体的组成部分,一般是基本数据类型
type 结构体名称 struct{
	field type
	field type
	...
}
~~~

## 创建结构体
~~~go
// 直接声明
var name struct
// {}
name := struct{}
// &
name := new(struct)
// 指针直接赋值 
&struct{val, val}
~~~




