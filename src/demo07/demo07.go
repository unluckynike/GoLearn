package main

import (
	"GoLearn/src/demo07/rectangle"
	"fmt"
	"log"
)

/*
包
*/

// 包级别变量
var rectLen, rectWidth float64 = -9, 7

// init 函数会检查长和宽是否大于0
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {

	/*
		1. import "packagename" 语句用于导入一个已存在的包 (GoLand软件可以自动导包)
	*/

	/*
				3. init 函数
			所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。
		init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性
		init 函数的形式如下：
			func init(){
			}

	*/

	//果然，程序会首先调用 rectangle 包的 init 函数，然后，会初始化包级别的变量 rectLen 和 rectWidth。
	//接着调用 main 包里的 init 函数，该函数检查 rectLen 和 rectWidth 是否小于 0，如果条件为真，则终止程序。
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))

	/*
						4.空白标识符
			导入了包，却不在代码中使用它，这在 Go 中是非法的。当这么做时，编译器是会报错的。其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加
		在程序开发的活跃阶段，又常常会先导入包，而暂不使用它。遇到这种情况就可以使用空白标识符 _。

				var _ = rectangle.Area // 错误屏蔽器
	*/
}
