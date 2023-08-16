package main

import "fmt"

/*
嵌套结构体
结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体。

*/

type Address struct {
	city  string
	state string
}

type person struct { //小写 私有
	name    string
	age     int
	address Address
}

type Staf struct {
	name    string
	age     int
	Address //匿名字段 Address，而 Address 是一个结构体。现在结构体 Address 有 city 和 state 两个字段，访问这两个字段就像在 Person 里直接声明的一样，因此我们称之为提升字段。
}

func main() {

	var p person
	p.name = "lisi"
	p.age = 18
	p.address = Address{"shanghai", "china"} //address，而 address 也是结构体

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	/*

		/*
		提升字段（Promoted Fields）
		如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。
	*/
	fmt.Println("======提升字段（Promoted Fields）=====")
	var s Staf
	s.name = "zhangsan"
	s.age = 18
	s.Address = Address{"beijing", "china"}
	fmt.Println("Name:", s.name)
	fmt.Println("Age:", s.age)
	fmt.Println("City:", s.city)   //city is promoted field
	fmt.Println("State:", s.state) //state is promoted field
}
