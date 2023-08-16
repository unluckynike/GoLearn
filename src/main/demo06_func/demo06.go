package main

import "fmt"

/*
函数
函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出
*/

func main() {

	/*
		1.函数声明
		在 Go 语言中，关键词 func 开始，后面紧跟自定义的函数名 functionname (函数名)，
		参数列表定义在 ( 和 ) 之间，返回值的类型则定义在之后的 returntype (返回值类型)处
		函数声明通用语法如下：
		func functionname(parametername type) returntype {
		    // 函数体（具体实现的功能）
		}
	*/
	function1()

	price := calculateBill(10, 5) //接收函数的total返回值
	fmt.Println("total price is : ", price)

	/*
		2.多返回值
	*/
	area, perimeter := rectProps(5, 5)
	fmt.Printf("area is :%d, perometer is :%d \n", area, perimeter)

	/*
		3.命名返回值
			从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
	*/
	fmt.Println(rectProps2(5, 5))

	/*
			4.空白符
		_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
	*/
	//以 rectProps 函数为例，该函数计算的是面积和周长。假使我们只需要计算面积，而并不关心周长的计算结果，空白符 _ 就上场了。
	a, _ := rectProps(5, 5) //只用到了函数 rectProps 的一个返回值 area 空白符 _ 用来跳过不要的计算结果。
	fmt.Println("只需要面积: ", a)

	/*
		5.可变参数
		...
	*/
	getSum(1, 2, 3, 4)
	getSum(1, 1, 1, 1, 1, 1, 1, 1)

}

func function1() {
	fmt.Println("函数声明，无参函数，无返回值函数")
}

// 有参数 有返回值 ,计算商品价格
func calculateBill(price int, num int) int {
	var total = price * num
	return total
}

// 多返回值函数 ,计算面积 周长
func rectProps(length int, width int) (int, int) { //括号里是返回值的类型，每个都用声明 ,如果一个函数有多个返回值，那么这些返回值必须用 ( 和 ) 括起来
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

// 命名返回值
func rectProps2(length, width int) (area, perimeter int) { //
	area = length * width
	perimeter = 2 * (length + width)
	return //函数中的 return 语句没有显式返回任何值。由于 area 和 perimeter 在函数声明中指定为返回值, 因此当遇到 return 语句时, 它们将自动从函数返回。
}

// 可变参数 函数参数可以有多个值
func getSum(nums ...int) { //可变参数必须放在参数最后一个
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
		sum += nums[i]
	}
	fmt.Println("sum: ", sum)
}
