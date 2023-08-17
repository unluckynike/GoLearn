package main

import "fmt"

/*
接口二
*/
type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

/*
实现多个接口
*/
type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() { //Employee 实现了LeaveCalculator接口
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int { //Employee 实现了CalculateLeavesLeft接口
	return e.totalLeaves - e.leavesTaken
}

/*
接口嵌套
Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口
*/
type EmployeeOperations interface {
	//接口 EmployeeOperations，它嵌套了两个接口：SalaryCalculator 和 LeaveCalculator
	SalaryCalculator
	LeaveCalculator
}

func main() {

	/*
		实现接口：指针接受者与值接受者
	*/
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/*
		对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。但接口中存储的具体值（Concrete Value）并不能取到地址，
		如果下面一行取消注释会导致编译错误：
		   cannot use a (type Address) as type Describer
		   in assignment: Address does not implement
		   Describer (Describe method has pointer
		   receiver)
	*/
	//d2 = a

	d2 = &a // 这是合法的 Address 类型的指针实现了 Describer 接口
	d2.Describe()

	/*
		实现多个接口
	*/
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	/*
		接口嵌套
	*/
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

	/*
		接口的零值
		接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil
	*/
	var d3 Describer
	if d3 == nil {
		fmt.Printf("d3 is nil and has type %T value %v\n", d3, d3)
	}
	//对于值为 nil 的接口，由于没有底层值和具体类型，当调用它的方法时，程序会产生 panic 异常
}
