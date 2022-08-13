# 单元测试
> Go语言中自带有一个轻量级的框架testing和自带的 `go test` 命令来实现测试和性能测试。

- 注意
  - 测试用例的文件名字以_test结尾.函数以Test开头
  - t.Fatalf格式化错误日志
  - t.Logf输出相应的日志
  - Pass表示Case成功, Fail表示失败
  - 测试单个文件 `go test -v Filename 源文件`
  - 测试单个方法 `go test -v -test.run MethodsName`
    

