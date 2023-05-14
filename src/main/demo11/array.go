package main

import "fmt"

/*
数组和切片
*/
func main() {

	/*
		1.数组 数组声明
		数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。
		Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。（译者注：如果是 interface{} 类型数组，可以包含任意类型）
		一个数组的表示形式为 [n]T。n 表示数组中元素的数量，T 代表每个元素的类型。元素的数量 n 也是该类型的一部分
		可以使用不同的方式来声明数组，让我们一个一个的来看。
	*/
	var a [3]int //初始化一个长度位3的数组a 数组中的所有元素都被自动赋值为数组类型的零值
	fmt.Println(a)

	var b [3]int
	b[0] = 2 //下标从0开始
	b[1] = 3
	b[2] = 5
	fmt.Println(b)

	//简略声明 来创建数组
	c := [3]int{4, 7, 9}
	fmt.Println(c)

	//不需要将数组中所有的元素赋值
	d := [4]int{23, 3} //声明了长度为4的数组d 但是只有前两个元素有值
	fmt.Println(d)

	//忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
	e := [...]int{23, 56, 78, 20}
	fmt.Println(e)

	//数组的大小是类型的一部分。因此 [5]int 和 [25]int 是不同类型。数组不能调整大小
	// a := [3]int{5, 78, 8}
	//    var b [5]int
	//    b = a // not possible since [3]int and [5]int are distinct types

	/*
		2.数组是值类型
		Go 中的数组是值类型而不是引用类型。意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，则不会影响原始数组。
	*/
	f := [...]string{"USA", "China", "India", "Germany", "France"}
	g := f //f赋值给g
	g[0] = "Singapore"
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	/*
		3.数组的长度
	*/
	h := [...]int{0, 1, 23, 232, 3, 434, 34, 3, 4}
	fmt.Println("h len: ", len(h))

	/*
			4.使用 range 迭代数组
		for 循环可用于遍历数组中的元素。
	*/
	arr := [...]float64{21, 2.2, 56.3, 67.5, 33.3}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d th element of arr is %.2f \n", i, arr[i])
	}

	//Go 提供了一种更好、更简洁的方法，通过使用 for 循环的 range 方法来遍历数组。
	//range 返回索引和该索引处的值。使用 range 重写上面的代码。还可以获取数组中所有元素的总和。
	fmt.Println("================range=================================")
	sum := float64(0)
	for i, v := range arr { //for i, v := range a 利用的是 for 循环 range 方式。 它将返回索引和该索引处的值
		fmt.Printf("%d th element of arr is %.2f \n", i, v)
		sum += v
	}
	fmt.Println("sum: ", sum)

	//如果只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行。
	for _, v := range arr { // ignores index
		fmt.Println(v)
	}
	/*
			5.多维数组
		Go 语言可以创建多维数组。
	*/
	array := [3][2]string{ //三行两列数组
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(array)
	var b_array [3][2]string
	b_array[0][0] = "apple"
	b_array[0][1] = "samsung"
	b_array[1][0] = "microsoft"
	b_array[1][1] = "google"
	b_array[2][0] = "AT&T"
	b_array[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b_array)

}

// 当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
