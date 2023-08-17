package main

import "fmt"

/*
可变参数函数
*/
func main() {
	/*
		语法
		如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数。
		只有函数的最后一个参数才允许是可变的。
	*/
	find(5, 3, 4, 5, 6, 7, 4, 0)
	find(2, 7, 8, 5, 5, 75, 6)

	/*
		给可变参数函数传入切片
	*/
	nums := []int{89, 90, 95}
	//find(89, nums) //无法通过编译，编译器报出错误
	find(89, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...) //使用了语法糖 ... 并且将切片作为可变参数传入 change 函数。使用了 ... ，welcome 切片本身会作为参数直接传入，不需要再创建一个新的切片。这样参数 welcome 将作为参数传入 change 函数
	fmt.Println(welcome)
}

func find(num int, nums ...int) { //最后一个是可变长参数
	fmt.Printf("类型: %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println("下标:", i)
			found = true
		}
	}
	if !found {
		fmt.Println("没找到")
	}

}

func change(s ...string) {
	s[0] = "Go"
}
