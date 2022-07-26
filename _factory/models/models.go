package models

// 定义私有的message
type message struct {
	Username, Password string
	age                int
}

// 添加值
func CreateUser(username, password string, age int) *message {
	return &message{
		Username: username,
		Password: password,
		age:      age,
	}
}

// 获取age属性
func (msg *message) GetMessage() int {
	return msg.age
}
