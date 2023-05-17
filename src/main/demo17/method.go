package main

import (
	"fmt"
	"math"
)

/*
方法
*/
type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() { //Employee 结构体的方法
	//在 Employee 结构体类型上创建了一个 displaySalary 方法。displaySalary()方法在方法的内部访问了接收器 e Employee
	fmt.Printf("salary of %s is %s%d \n", e.name, e.currency, e.salary)
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	/*
		方法
		方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
		下面就是创建一个方法的语法。
		func (t Type) methodName(parameter list) {
		}
	*/
	emp1 := Employee{
		name:     "sam Adolf",
		salary:   888,
		currency: "$",
	}
	emp1.displaySalary()

	//why? Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径
	displaySalary(emp1)
	fmt.Println()

	r := Rectangle{length: 10, width: 10}
	fmt.Println("矩形面积： ", r.Area())

	c := Circle{radius: 5}
	fmt.Println("圆面积： ", c.Area())

	/*
		指针接收器与值接收器
	*/
}
