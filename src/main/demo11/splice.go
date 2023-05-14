package main

import "fmt"

/*
切片
*/
func main() {

	/*
		1.切片
		切片("动态数组")是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
		与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
		声明一个未指定大小的数组来定义切片：
		var identifier []type
		切片不需要说明长度。
		或使用 make() 函数来创建切片:
		var slice1 []type = make([]type, len)
		也可以简写为
		slice1 := make([]type, len)
		也可以指定容量，其中 capacity 为可选参数。
		make([]T, length, capacity)
		这里 len 是数组的长度并且也是切片的初始长度。
	*/
	//带有 T 类型元素的切片由 []T 表示
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] //使用语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片。
	fmt.Println(b)
	//另一种创建切片的方法
	c := []int{6, 7, 8}
	fmt.Println(c)

	/*
		2.切片的修改
	*/
}
