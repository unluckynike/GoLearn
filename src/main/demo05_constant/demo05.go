package main

/*
常量
*/
import (
	"fmt"
)

func main() {

	/*
		1.定义常量
	*/
	//没有明确的使用关键字 const，但是在 Go 的内部，它们是常量。
	var a int = 5
	var b string = "Learn Go "
	fmt.Println(a, b)

	//常量不能再重新赋值为其他的值
	const c = 20
	//c = 19 // 不允许重新赋值

	//常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。
	fmt.Println("Hello, playground")
	//var a = math.Sqrt(4)   // 允许
	//const b = math.Sqrt(4) // 不允许

	/*
		2. 字符串常量
		双引号中的任何值都是 Go 中的字符串常量
	*/

	const hello = "Hello World"             //Hello World 这样的字符串常量没有任何类型
	const typedhello string = "Hello World" //typedhello 就是一个 string 类型的常量

	//Go 是一个强类型的语言，在分配过程中混合类型是不允许的
	//var defaultName = "Sam" // 允许
	//type myString string
	//var customName myString = "Sam" // 允许
	//customName = defaultName // 不允许

	/*
		3. 布尔常量
		布尔常量和字符串常量没有什么不同。他们是两个无类型的常量 true 和 false。字符串常量的规则适用于布尔常量
	*/

	/*
		4.数字常量
		数字常量包含整数、浮点数和复数的常量
	*/
	const e = 5
	var intVar int = e
	var int32Var int32 = e
	var float64Var float64 = e
	var complex64Var complex64 = e
	//以代表一个浮点数、整数甚至是一个没有虚部的复数
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	//默认类型取决于语法 每个变量的类型由数字常量的语法决定
	var i = 5
	var f = 5.6
	var g = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, g's type %T\n", i, f, g)

	/*
			5.数字表达式
		数字常量可以在表达式中自由混合和匹配，只有当它们被分配给变量或者在需要类型的代码中的任何地方使用时，才需要类型。
	*/
	var h = 8.8 / 2
	fmt.Printf("type: %T, value: %v", h, h) //%v 是一种格式化输出方式，用于将变量的值以默认格式打印出来。%v用于任何类型的变量，包括字符串、整数、浮点数、数组、切片、结构体等，也可以用于自定义类型。
}
