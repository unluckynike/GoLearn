package main

import (
	"fmt"
	"unicode/utf8"
)

/*
字符串
*/

func main() {

	/*
		字符串
		Go 语言中的字符串是一个字节切片。把内容放在双引号""之间，兼容 Unicode 编码的，并且使用 UTF-8 进行编码。

	*/
	name := "Hello World"
	fmt.Println(name)

	//由于字符串是一个字节切片，所以我们可以获取字符串的每一个字节。
	for i := 0; i < len(name); i++ {
		//len(s) 返回字符串中字节的数量，然后我们用了一个 for 循环以 16 进制的形式打印这些字节。%x 格式限定符用于指定 16 进制编码。
		//程序输出 48 65 6c 6c 6f 20 57 6f 72 6c 64。这些打印出来的字符是 "Hello World" 以 Unicode UTF-8 编码的结果。
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	//打印字符串的每一个字符
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i]) //%c 格式限定符用于打印字符串的字符
	}

	/*
		rune
		rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
	*/
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)

	/*
			字符串的循环
		 Go 提供了一种更简单的方法来做到这一点：使用 for range 循环。
	*/
	for index, rune := range name { //使用 for range 循环遍历了字符串。循环返回的是是当前 rune 的字节位置。
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

	/*
		切片构造字符串
	*/
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072} //runeSlice 包含字符串 Señor的 16 进制的 Unicode 代码点。这个程序将会输出Señor。
	str := string(runeSlice)
	fmt.Println(str)

	/*
			字符串长度
		utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。
	*/
	fmt.Println(len(name), name)
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	/*
		字符串修改
			Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。
		为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。
	*/
	h := "hello"
	fmt.Println(mutate([]rune(h))) //传入的就是字符串切片

}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	//字符串被转化为一个 rune 切片。然后我们循环打印字符
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	//mutate 函数接收一个 rune 切片参数，它将切片的第一个元素修改为 'a'，然后将 rune 切片转化为字符串，并返回该字符串。
	s[0] = 'a'
	return string(s)
}
