## 字符串

### 常用方法
- 计算字符串长度 len()
- 十进制转其他进制 strconv.FormatInt()
- 字符串转整数 strconv.Atoi("123")
- 整数转字符串 strconv.Itoa(110)
- 判断子串是否在指定的字符串中 strings.Contains("cheese","e")
- 统计一个字符串中有几个子串 strings.Count("cheese","e")
- 不区分大小写的字符串比较 strings.EqualFold("ABc", "abc")
- 返回子串在字符串中第一次出现的位置,没有则返回-1 strings.Index("cheese","a")
- 返回子串在字符串中最后一次出现的位置,没有则返回-1 strings.astIndex("cheese","a")
- 字符串替换strings.replace(string, 换那个，换成什么，换几个)   -1 是所有
- 按照指定的字符进行替换strings.Split(string, ",") 返回切片
- 字符串大小写转换 ToLower()  ToUpper()
- 将字符串左右两侧的空格去掉 strings.TrimSpace()
- 去除字符串两边指定的字符 strings.Trim()
- 去除字符串左右两边 TrimLeft() TrimRight()
- 判断字符串是否以指定字符串开头 strings.HasPrefix("hello", "he")
- 判断字符串是否以指定字符串结尾 strings.HasSuffix("hello", "he")