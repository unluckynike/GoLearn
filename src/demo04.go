package main

import (
	"fmt"
	"unsafe"
)

/*
类型
*/

func main() {
	/*
		Go 支持的基本类型：

		-bool
		-数字类型
			int8, int16, int32, int64, int 整数类型，分别表示有符号整数
			uint8, uint16, uint32, uint64, uint 无符号整数类型
			float32, float64 浮点数类型，分别表示单精度浮点数和双精度浮点数
			complex64, complex128 复数类型，分别表示32位和64位的复数
			byte 一种整数类型，byte 类型是一个代表单个字节的无符号整数类型。它实际上是 uint8 的别名u
			rune 一种整数类型，它用于表示 Unicode 字符。具体来说，rune 类型实际上是 int32 类型的别名。
		-string
	*/

	/*
		1. bool
		bool 类型表示一个布尔值，值为 true 或者 false
	*/
	a, b := true, false
	fmt.Println("a:", a, "b:", b)
	c := a && b //仅当 a 和 b 都为 true 时，操作符 && 才返回 true
	fmt.Println("c:", c)
	d := a || b //当 a 或者 b 为 true 时，操作符 || 返回 true
	fmt.Println("d:", d)

	/*
		2. 有符号整型
		int8：表示 8 位有符号整型
		大小：8 位
		范围：-128～127

		int16：表示 16 位有符号整型
		大小：16 位
		范围：-32768～32767

		int32：表示 32 位有符号整型
		大小：32 位
		范围：-2147483648～2147483647

		int64：表示 64 位有符号整型
		大小：64 位
		范围：-9223372036854775808～9223372036854775807

		int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
		大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	*/
	var e int = 89
	f := 95
	fmt.Println("value of e is", e, "and f is", f)
	fmt.Printf("type of e is %T, size of e is %d", e, unsafe.Sizeof(e))   // e 的类型和大小
	fmt.Printf("\ntype of f is %T, size of f is %d", f, unsafe.Sizeof(f)) // f 的类型和大小
	//32位 大小都是 32 位（4 字节） ，64 位系统下，a 和 b 会占用 64 位（8 字节）的大小

	/*
		3. 无符号整型
		uint8：表示 8 位无符号整型
		大小：8 位
		范围：0～255

		uint16：表示 16 位无符号整型
		大小：16 位
		范围：0～65535

		uint32：表示 32 位无符号整型
		大小：32 位
		范围：0～4294967295

		uint64：表示 64 位无符号整型
		大小：64 位
		范围：0～18446744073709551615

		uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
		大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	*/

	var g uint = 9 // 无符号整型
	h := 10
	fmt.Printf("\n%T,%d", g, unsafe.Sizeof(g))
	fmt.Printf("\n%T,%d\n", h, unsafe.Sizeof(h))

	/*
		4. 浮点型
		float32：32 位浮点数
		float64：64 位浮点数
	*/

	i, j := 25.2, 10.5
	fmt.Printf("Type i %T,Type j %T \n", i, j)
	sum := i + j
	diff := i - j
	fmt.Println("sum: ", sum, "diff: ", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	/*
		5. 复数类型
		complex64：实部和虚部都是 float32 类型的的复数。
		complex128：实部和虚部都是 float64 类型的的复数。

		内建函数 complex 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：
		func complex(r, i FloatType) ComplexType 该函数的参数分别是实部和虚部，并返回一个复数类型。实部和虚部应该是相同类型

		使用简短语法来创建复数：
		c := 6 + 7i
	*/

	c1 := complex(4, 7)
	c2 := complex(5, 3)
	cadd := c1 + c2
	fmt.Println("complex sum: ", cadd)
	cmul := c1 * c2 //c1 和 c2 的乘积赋值给 cmul
	//设z1=a+bi，z2=c+di(a、b、c、d∈R)是任意两个复数，那么它们的积(a+bi)(c+di)=( ac-bd)+(bc+ad)i
	fmt.Println("product: ", cmul)

	//简短语法创建
	c3 := 3 + 2i
	fmt.Println(c3)

	/*
		6. 其他数字类型
		byte 是 uint8 的别名。
		rune 是 int32 的别名。
	*/

	/*
		7. string 类型
		字符串是字节的集合
	*/

	first := "Alice"
	Last := "Park"
	name := first + Last
	fmt.Println("my name is: ", name)

	/*
		8. 类型转换
		Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换
	*/

	k, l := 55, 25.5
	//fmt.Println("sum: ",k+l)//不允许int + float
	fmt.Println("sum: ", float64(k)+l) //强转类型

	//赋值的情况也是如此。把一个变量赋值给另一个不同类型的变量，需要显式的类型转换
	m := 10
	var n float64 = float64(m) // 若没有显式转换，该语句会报错
	fmt.Println("n:", n)
}
