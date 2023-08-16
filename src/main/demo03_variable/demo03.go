package main

/*
go 变量
*/

import (
	"fmt"
	"math"
)

func main() {
	/*
		1.声明单个变量
	*/
	//var name type 是声明单个变量的语法。
	var age int                   //变量声明
	fmt.Println("my age is", age) //默认值是 0

	//赋值
	age = 22
	fmt.Println("my age is", age)

	//age 可以赋值为任何整型值（Integer Value）
	age = 56
	fmt.Println(age)

	/*
		2.声明变量并做初始化
	*/
	var name string = "Lin"
	fmt.Println("my name is :", name)

	/*
		3.类型推断
	*/
	//如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。
	//因此，如果变量有初始值，就可以在变量声明中省略 type

	var num = 22 // 可以推断类型 var name = initialvalue，Go 能够根据初始值自动推断变量的类型
	fmt.Println(num)

	/*
		4.声明多个变量
		声明多个变量的语法是 var name1, name2 type = initialvalue1, initialvalue2
	*/
	var width, height int = 100, 50 //int 也可以省略
	fmt.Println(width, height)

	//在一个语句中声明不同类型的变量
	var (
		id       = 1
		username = "Zhang"
		password = "abc123"
	)
	fmt.Println(id, username, password)

	/*
		5 简短声明
		Go 也支持一种声明变量的简洁形式，称为简短声明（Short Hand Declaration），该声明使用了 := 操作符。
		声明变量的简短语法是 name := initialvalue
	*/
	sex, birth := "男", 29
	fmt.Println(sex, birth)

	//x,y := 1 //error y没有被赋值
	//fmt.Println(x,y)

	//简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的
	a, b := 20, 30 // 声明变量a和b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("changed b is", b, "c is", c)

	//a, b := 20, 30 // 声明a和b
	//fmt.Println("a is", a, "b is", b)
	//a, b := 40, 50 // 错误，没有尚未声明的变量  no new variables on left side of := 的错误，这是因为 a 和 b 的变量已经声明过了，:= 的左边并没有尚未声明的变量

	fmt.Println("Max: ", math.Max(float64(a), float64(b)))

	Alice, number, flag, flotNum := "Alice", 12, false, 12.2
	fmt.Printf("%T,%T,%T,%T", Alice, number, flag, flotNum) //输出类型

}
