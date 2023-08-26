package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"proto-demo/proto/userService"
)

func main() {
	user := &userService.Userinfo{
		Username: "light",
		Age:      18,
		Hobby:    []string{"吃饭", "睡觉"},
	}
	// Protobuf序列化
	data, _ := proto.Marshal(user)
	fmt.Println("序列化后的数据:", data)
	// Protobuf反序列化
	u := userService.Userinfo{}
	proto.Unmarshal(data, &u)
	fmt.Printf("反序列化后的数据：%#v", user)
}
