package main

import (
	"fmt"
	"main/golangTutorialSeries/demo07/rectangle"
)

func main() {
	/*
		2. 自定义包 导入
		rectangle 包中的函数 Area 和 Diagonal 首字母大写。在 Go 中这具有特殊意义。在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。
		在这里，需要在 main 包中访问 Area 和 Diagonal 函数，因此会将它们的首字母大写。
	*/
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))

}
