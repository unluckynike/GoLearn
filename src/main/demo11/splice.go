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
	var b []int = a[1:4] //使用语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片。 开始和结束的默认值分别为 0 和 len (numa)
	fmt.Println(b)
	//另一种创建切片的方法
	c := []int{6, 7, 8}
	fmt.Println(c)

	var numbers = make([]int, 3, 5) //通过make 创建切片
	printSlice(numbers)

	/*
			2.切片的修改
		切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
	*/
	darr := [...]int{2, 3, 4, 5, 5, 3, 3, 4, 3, 432, 23, 56}
	dslice := darr[2:5] //取到下标 2到5到元素
	fmt.Println("darr: ", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("darr: ", darr)

	//当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
	//切片共享同一个数组时，每个所做的修改都会反映在数组中。
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //取到所有元素
	nums2 := numa[:]
	fmt.Println("array before change 1", numa) //78，79，80
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa) //100 79 80
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa) //100 101 80

	/*
		3.切片的长度和容量
		切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
	*/

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3] //"orange", "grape"
	//fruitslice 的容量是从 fruitarray 索引为 1 开始，也就是说从 orange 开始，该值是 6
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6

	fruitslice = fruitslice[:cap(fruitslice)] // re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	/*
			4.追加元素
		数组的长度是固定的，它的长度不能增加。
		切片是动态的，使用 append 可以将新元素追加到切片上。
		append 函数的定义是 func append（s[]T，x ... T）[]T
	*/
	//当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用。现在新切片的容量是旧切片的两倍。

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6

	//可以使用 ... 运算符将一个切片添加到另一个切片
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...) //可变参数 追加切片
	fmt.Println("food:", food)

	/*
				5.切片的函数传递
			切片在内部可由一个结构体类型表示。
			它的表现形式
			type slice struct {
			    Length        int 长度
			    Capacity      int 容量
			    ZerothElement *byte 第零个元素的指针
			}
		当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。
	*/

	nos := []int{8, 6, 6, 6}
	fmt.Println("nos: ", nos)
	subtactOne(nos)
	fmt.Println("nos: ", nos)

	/*
		多维切片
	*/

	pls := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{2, 3, 4},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Print(" ", v2)
		}
		fmt.Println()
	}

	/*
		内存优化
		切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。
		假设有一个非常大的数组，我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。
		一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
	*/
	countriesNeeded := countries() //将 neededCountries 复制到 countriesCpy 同时在函数的下一行返回 countriesCpy。现在 countries 数组可以被垃圾回收, 因为 neededCountries 不再被引用。
	fmt.Println(countriesNeeded)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func subtactOne(numbers []int) {
	for i, _ := range numbers {
		numbers[i]++
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
