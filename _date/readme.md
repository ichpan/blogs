# Golang时间日期函数

### time方法使用

- 获取当前时间
~~~go
time.Now()
~~~

- 获取当前的时分秒
~~~go
now := time.Now()
// 获取年份
now.Year()
// 获取月份
now.Month()
// 获取日
now.Day()
// 获取小时
now.Hour()
// 获取分钟
now.Minute()
// 获取秒数
now.Second()
~~~

- 休眠
~~~go
// 休眠一秒
time.Sleep(time.Second)

// 休眠一毫秒
time.Sleep(time.Millisecond * 100)
~~~
- 时间戳
~~~go
// 秒为单位
now.Unix()

// 纳秒为单位
now.UnixNano
~~~