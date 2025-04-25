package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 模拟从外部传入 interface{}
	var i interface{} = User{Name: "Yitan", Age: 22}

	user, ok := i.(User)
	if ok {
		fmt.Printf("断言成功: %+v\n", user)
		fmt.Printf("%v\n", user) // 输出: {Yitan 22}
	} else {
		fmt.Println("断言失败")
	}

	var i2 interface{} = "not a user"

	user2, ok := i2.(User)
	if !ok {
		fmt.Println("这不是 User 类型")
	} else {
		fmt.Printf("断言成功: %+v\n", user2)
	}

}
