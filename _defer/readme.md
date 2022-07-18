## defer使用
**为什么需要defer？**

&emsp;&emsp;我们需要经常创建资源(问件句柄,数据库链接,锁等),为了在函数执行完毕后,及时**释放资源**,**延迟**释放。

- 注意:当函数执行到defer语句的时候,会将defer后面的语句压入一个独立的栈中。
- 栈 Lifo原则

值传递和引用传递
- 值传递是值的拷贝
- 引用传递是地址的拷贝
- 地址拷贝效率高

值类型: 基本数据类型 int、float、bool、string、array、struct...
引用类型: 指针、slice、map、channel、interface...
