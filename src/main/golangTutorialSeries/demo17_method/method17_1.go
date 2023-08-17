package main

import "fmt"

type employee struct {
	name string
	age  int
}

/*
使用值接收器的方法。
定义了一个名为changeName的方法，该方法属于employee结构体。
changeName方法的作用是将employee的name属性更改为传入的newName参数。
*/
func (e employee) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *employee) changeAge(newAge int) {
	e.age = newAge
}

/*
匿名字段的方法
*/

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

/*
别名
定义了一个名为myInt的新类型，它是int的别名。
为myInt类型定义了一个方法add，该方法接受一个myInt类型的参数b，并返回两个参数的和。
*/

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	/*
		指针接收器与值接收器
		值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。
	*/
	e := employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew") //changeName 方法中对 Employee 结构体的字段 name 所做的改变对调用者是不可见的，因此程序在调用 e.changeName("Michael Andrew") 这个方法的前后打印出相同的名字
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51) //changeAge 方法是使用指针 (e *Employee) 接收器的，所以在调用 (&e).changeAge(51) 方法对 age 字段做出的改变对调用者将是可见的
	fmt.Printf("\nEmployee age after change: %d", e.age)
	e.changeAge(55) //另外写法
	fmt.Printf("\nEmployee age after change: %d", e.age)
	fmt.Println()
	/*
		一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。

		指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。

		在其他的所有情况，值接收器都可以被使用。
	*/

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	//通过使用 p.fullAddress() 来访问 address 结构体的 fullAddress() 方法
	p.fullAddress() //访问 address 结构体的 fullAddress 方法
	fmt.Println()

	/*
		当一个函数有一个值参数，它只能接受一个值参数。
		当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	*/

	/*
		在方法中使用指针接收器 与 在函数中使用指针参数
	*/

	/*
			别名
		为 int 创建了一个类型别名 myInt。定义一个以 myInt 为接收器的的方法 add。
	*/
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
