package main

import "fmt"

/*
指针
*/
func main() {
	/*
		指针
		指针是一种存储变量内存地址（Memory Address）的变量。
		Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
		指针变量的类型为 *T，该指针指向一个 T 类型的变量。
	*/
	b := 255
	var a *int = &b //把 b 的地址赋值给 *int 类型的 a。称 a 指向了 b。当打印 a 的值时，会打印出 b 的地址
	fmt.Printf("a的类型 %T\n", a)
	fmt.Printf("b的地址 ", a)
	fmt.Println()

	/*
		指针的零值
		指针的零值是 nil
	*/
	var c *int
	if c == nil {
		fmt.Println("c is： ", c) //c指针零值
		c = &b
		fmt.Println("c after initialization is: ", c)
	}

	/*
		数组指针
	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1的地址: %p\n", &arr1)
	var p1 *[4]int
	p1 = &arr1

	fmt.Printf("p1->地址: %p\n", p1)
	fmt.Printf("p1地址: %p\n", &p1)
	fmt.Println("p1->地址->值： ", *p1)

	/*
		解引用
		指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a。
	*/
	d := 255
	e := &d
	fmt.Println("address of d is: ", e) //地址
	fmt.Println("value of d is: ", *e)  //值

	*e++ //修改d的值
	fmt.Println("new value d is: ", d)

	/*
		向函数传递指针参数
	*/

	f := 58
	fmt.Println("f: ", f)
	g := &f
	change(g)
	fmt.Println("f: ", f)

	//不要向函数传递数组的指针，而应该使用切片
	//假如我们想要在函数内修改一个数组，并希望调用函数的地方也能得到修改后的数组，一种解决方案是把一个指向数组的指针传递给这个函数。
	h := [3]int{98, 78, 67}
	modify(&h)
	fmt.Println("h: ", h)

	//用切片重构
	i := [3]int{89, 90, 91}
	modify_splice(i[:])
	fmt.Println(i)

	////Go 不支持指针运算
	//b := [...]int{109, 110, 111}
	//p := &b
	//p++ //会抛出编译错误
}

func change(val *int) { //向函数 change 传递了指针变量 b，而 b 存储了 a 的地址。在 change 函数内使用解引用，修改了 a 的值。
	*val = 55
}

func modify(arr *[3]int) { //将数组的地址传递给了 modify 函数。在第 8 行，在 modify 函数里把 arr 解引用，并将 90 赋值给这个数组的第一个元素
	(*arr)[0] = 90 ////a[x] 是 (*a)[x] 的简写形式，因此代码中的 (*arr)[0] 可以替换为 arr[0]
}

// 别再传递数组指针了，而是使用切片吧。代码更加简洁，也更符合 Go 语言的习惯。
func modify_splice(sls []int) { //将一个切片传递给了 modify 函数。在 modify 函数中，我们把切片的第一个元素修改为 90
	sls[0] = 90
}
