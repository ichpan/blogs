## 管道`Channel`

> **引用类型**,make之后才可以使用

### 基本介绍

- Channel本质就是一个**队列**，遵循**FIFO**原则。

- Channel本身就是线程安全的。多goroutine访问不需要加锁。

- Channel是有类型的。

### 基本用法

~~~go
// 声明语法
var 变量名 chan 数据类型

// 例子
mychan := make(chan int, 3)

// 写入数据  不可以超过容量
mychan <- 5

// 取出数据
num := <- mychan

// 计算容量/大小
len(mychan)  cap(mychan)
~~~

### 关闭和便利

~~~go
// 使用内置的方法close可疑关闭管道
channel.close()

// 便利
for v := range mychan{
    fmt.Println(v)
}
~~~

### 使用细节

- channel只可以存放指定的数据类型
- channel的数据放满后就不可以再放入了
- 放满数据后再取出数据可以继续放入
- 没有使用goroutine的情况下，如果数据取完后再取数据，就会报dead lock
- 可以不接收从管道中取出来的数据

- 管道关闭之后，不可以写入数据，只能取数据
- 便利管道只能使用range方法
- 在便利管道的时候，如果管道没有关闭会产生死锁
- 如果只向管道中写入数据，而没有读取，就会出现阻塞，不需要考虑读写速度的问题
- 管道可以声明为只读或者只写，默认是双向的
- 使用select可以解决从管道中取数据阻塞问题
- goroutine中使用recover，解决panic，导致程序崩溃问题

