package main

import "fmt"

// map splice 综合案例
func main() {

	user1 := make(map[string]string)
	user1["name"] = "park"
	user1["age"] = "20"
	user1["sex"] = "男"
	user1["addr"] = "上海"

	user2 := make(map[string]string)
	user2["name"] = "alice"
	user2["age"] = "23"
	user2["sex"] = "男"
	user2["addr"] = "北京"

	user3 := map[string]string{ //不用make 初始化
		"name": "jack",
		"age":  "19",
		"sex":  "女",
		"addr": "深圳",
	}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
	//定义的切片切片长度为 0，容量为 3  元素类型是map类型 map里键值对是 stirng string类型，
	userData := make([]map[string]string, 0, 3)
	userData = append(userData, user1)
	userData = append(userData, user2)
	userData = append(userData, user3)

	fmt.Println(userData)
}
