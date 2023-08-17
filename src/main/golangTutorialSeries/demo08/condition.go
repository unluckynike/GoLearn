package main

import "fmt"

/*
if-else
*/
func main() {

	/*
		1. if-else
		if 是条件语句。if 语句的语法是

		if condition {
		}
		如果 condition 为真，则执行 { 和 } 之间的代码。

		不同于其他语言，例如 C 语言，Go 语言里的 { } 是必要的，即使在 { } 之间只有一条语句。

		if 语句还有可选的 else if 和 else 部分。

		if condition {
		} else if condition {
		} else {
		}
		if-else 语句之间可以有任意数量的 else if。条件判断顺序是从上到下。如果 if 或 else if 条件判断的结果为真，则执行相应的代码块。 如果没有条件为真，则 else 代码块被执行
	*/

	num := 10
	if num%2 == 0 {
		fmt.Println("even")
	} else { //else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。 原因：Go 语言的分号是自动插入指定在 } 之后插入一个分号，如果这是该行的最终标记
		fmt.Println("odd")
	}
	/*
		2.其他形式

		if 还有另外一种形式，它包含一个 statement 可选语句部分，该组件在条件判断之前运行。它的语法是

		if statement; condition {
		}
	*/

	if n := 10; n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
