package main

import "fmt"

/*
switch语句
*/
func main() {
	/*
			1. switch
		switch 是一个条件语句，用于将表达式的值与可能匹配的选项列表进行比较，并根据匹配情况执行相应的代码块。它可以被认为是替代多个 if else 子句的常用方式
	*/
	finger := 4
	switch finger {
	case 1: //case 不允许出现重复项
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // 默认情况 当其他情况都不匹配时，将运行默认情况。
		fmt.Println("incorrect finger number")
	}

	/*
			2. 多表达式判断
		通过用逗号分隔，可以在一个 case 中包含多个表达式。
	*/
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //一个选项多个表达式 在 case "a","e","i","o","u": 这一行中，列举了所有的元音。只要匹配该项，则将输出 vowel。
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	/*
			3.无表达式的switch
		在 switch 语句中，表达式是可选的，可以被省略。如果省略表达式，则表示这个 switch 语句等同于 switch true，并且每个 case 表达式都被认定为有效，相应的代码块也会被执行。
	*/
	num := 75
	switch { // 表达式被省略了 默认它为 true，true 值会和每一个 case 的求值结果进行匹配
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	/*
			4.Fallthrough 语句
		在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	*/
	switch num := number(); { // num is not a constant , switch 和 case 的表达式不一定是常量
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough //fallthrough 语句应该是 case 子句的最后一个语句。
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
	//还有一种 switch 类型称为 type switch 。接口的时候再研究这个。
}

func number() int {
	num := 15 * 5
	return num
}
