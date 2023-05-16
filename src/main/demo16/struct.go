package main

import "fmt"

/*
结构体
*/
//Employee 权限 小写私有，大写公开
type Employee struct {
	firstName, lastName string
	age, salary         int
}

type User struct {
	name string
	age  int
	sex  string
}

// 匿名字段 创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。
type Person struct {
	string
	int
}

func main() {
	/*
		结构体
		结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
		结构体声明
		type Employee struct {
		    firstName string
		    lastName  string
		    age       int
		}
		声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 匿名结构体（Anonymous Structure）。
		var employee struct {
		    firstName, lastName string
		    age int
		}
	*/
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800} //省略了字段名。在这种情况下，就需要保证字段名的顺序与声明结构体时的顺序相同。

	var emp3 Employee //创建对象
	emp3.firstName = "alice"
	emp3.age = 30
	emp3.salary = 800
	emp3.lastName = "Liu"

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)

	/*
		创建匿名结构体
	*/
	emp4 := struct { //定义了一个匿名结构体变量 emp3。之所以称这种结构体是匿名的，是因为它只是创建一个新的结构体变量 em3，而没有定义任何结构体类型。
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 4", emp4)

	/*
		结构体相等性
		结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
		结构体指针应用
	*/
	//结构体是值类型的，值传递
	user1 := User{"张三", 18, "男"}
	fmt.Println(user1)
	user2 := user1
	//创建零值的 struct，以后再给各个字段赋值。
	user2.name = "李四"
	fmt.Println(user1) //不影响原来的值
	fmt.Println(user2)

	/*
		结构体的零值
	*/

	var emp5 Employee           //零值结构体
	fmt.Println("emp5: ", emp5) //firstName 和 lastName 赋值为 string 的零值（""）。而 age 和 salary 赋值为 int 的零值（0）

	//部分指定初始值
	emp7 := Employee{
		firstName: "zhang",
		lastName:  "sun",
	}
	fmt.Println("emp6: ", emp7)

	//访问结构体的字段 点号操作符 . 用于访问结构体的字段。
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d \n", emp6.salary)

	/*
		结构体指针
	*/
	emp8 := &Employee{"Sam", "Anderson", 55, 900}
	fmt.Println("First Name: ", (*emp8).firstName) //emp8 是一个指向结构体 Employee 的指针。(*emp8).firstName 表示访问结构体 emp8 的 firstName 字段
	//可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName。
	fmt.Println("first name: ", emp8.firstName)
	fmt.Println("Age: ", (*emp8).age)

	/*
			匿名字段
		虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。
		比如在上面的 Person 结构体里，虽说字段是匿名的，但 Go 默认这些字段名是它们各自的类型。所以 Person 结构体有两个名为 string 和 int 的字段。
	*/
	p := Person{"Naveen", 50}
	fmt.Println("p: ", p)

	/*
		嵌套结构体
	*/
}
