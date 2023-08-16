package main

import "fmt"

/*
接口
*/

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// 实现接口
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

/*
接口的内部表示
*/

type Test interface {
	Tester()
}
type MyFloat float64

func (m MyFloat) Tester() { //MyFloat类型实现了 Tester 接口
	fmt.Println(m)
}
func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

/*
空接口
*/
func describe2(i interface{}) { //函数接收空接口作为参数，因此，可以给这个函数传递任何类型。
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

/*
接口断言
*/

func assert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}

func assert2(i interface{}) {
	/*
		如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
		如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。
	*/
	v, ok := i.(int)
	fmt.Println(v, ok)
}

/*
类型选择
*/
func findType(i interface{}) {
	switch i.(type) { //switch i.(type) 表示一个类型选择。每个 case 语句都把 i 的具体类型和一个指定类型进行了比较。如果 case 匹配成功，会打印出相应的语句
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	case float64:
		fmt.Printf("I am an float64 and my value is %f\n", i.(float64))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func main() {

	/*
		接口
		在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。
		在 Go 语言中，接口就是方法签名（Method Signature）的集合。当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。
		接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
		如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口。
	*/

	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("vowels are %c", v.FindVowels())

	/*
		接口的内部表示
		可以把接口看作内部的一个元组 (type, value)。 type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
	*/
	var t Test
	f := MyFloat(89.03)
	t = f
	describe(t)
	t.Tester()

	/*
		空接口
		没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。
	*/
	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)

	/*
		类型断言
		类型断言用于提取接口的底层值（Underlying Value）。
		在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。
	*/
	var ss interface{} = 55
	assert(ss)
	//var sss interface{} = "dfdf" //运行报错
	//assert(sss)

	var sa interface{} = 56
	assert2(sa)
	var ia interface{} = "Steven Paul"
	assert2(ia)

	/*
		类型选择
		类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。
		类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，而对于类型选择，类型 T 由关键字 type 代替。
	*/
	findType("this is string")
	findType(532)
	findType(56.9)
}
