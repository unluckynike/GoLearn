package rectangle

import (
	"fmt"
	"math"
)

/*
 * init function added
 */
func init() {
	fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) (area float64) { //函数挎包对外调用 方法名首字母大写
	area = len * wid
	return
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
