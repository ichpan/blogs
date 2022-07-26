package main

import (
	"blogs/_factory/models"
	"fmt"
)

func main() {
	user := models.CreateUser("zhangsan", "qwetr", 18)
	fmt.Println(user.Username, user.Password, user.GetMessage())
}
