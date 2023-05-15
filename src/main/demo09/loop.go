package main

import "fmt"

/*
循环
*/

func main() {
	/*
		1. 循环语法
		for 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环。
		for 循环语法
		for initialisation; condition; post {
		}
		三个组成部分，即初始化，条件和 post 都是可选的。让我们看一个例子来更好地理解循环。
	*/
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println("")

	/*
		2. break
		break 语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。
	*/

	for j := 0; j < 10; j++ {
		if j > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf(" %d", j)
	}
	fmt.Println("\n line after for loop ")

	/*
		3. continue
		continue 语句用来跳出 for 循环中当前循环。在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。循环体会在一下次循环中继续执行。
	*/

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //偶数就跳过
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println(" ")

	/*
		4. 其他例子
	*/
	i := 0
	for i < 10 { //for 循环的三部分，初始化语句、条件语句、post 语句都是可选的
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println(" ")

	//在 for 循环中可以声明和操作多个变量
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	/*
		5. 无限循环
			无限循环的语法是：
			for {
			}
	*/
	//for {
	//	fmt.Println("Learn Go")
	//}

	/*
			6. 嵌套循环
		 Go 语言嵌套循环的格式：

		for [condition |  ( init; condition; increment ) | Range]
		{
		   for [condition |  ( init; condition; increment ) | Range]
		   {
		      statement(s);
		   }
		   statement(s);
		}
	*/

	var k, l int
	for k = 2; k < 100; k++ {
		for l = 2; l <= (k / l); l++ {
			if k%l == 0 {
				break //发现因子不是素数
			}
		}
		if l > (k / l) {
			fmt.Printf("%d ", k)
		}
	}

	/*
		range
		for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

		for key, value := range oldMap {
		    newMap[key] = value
		}
	*/
}
