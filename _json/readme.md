## GOLANG处理JSON

### 序列化
~~~go
// encoding/json 模块
json.Marshal()
~~~

### 反序列化
~~~go
// encoding/json 模块
json.Unmarshal([]byte(data), &struct)
~~~

### tag标签的使用
~~~go
// 采用反射机制
`json:"newname"`
~~~