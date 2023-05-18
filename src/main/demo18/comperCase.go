package main

import "fmt"

// 将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func find(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Println("unknow type \n")
	}
}
func main() {

	find("alice")
	find(123)
	find(55.3)

	p := Person{
		name: "Naveen R",
		age:  18,
	}
	find(p) //Person类

}
