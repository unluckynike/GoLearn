package main

import (
	"fmt"
	"main/golangTutorialSeries/demo16_struct/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	//spec.model = "mac mini" //首字母小写了 是私有的 如果我们试图访问未导出的字段 model，编译器会报错
	spec.Price = 9000
	fmt.Println("spec: ", spec)
}
