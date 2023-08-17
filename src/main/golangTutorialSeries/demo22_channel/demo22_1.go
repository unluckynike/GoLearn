package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	//time.Sleep(1 * time.Second)
	fmt.Println("main function")

	//在一个单独的 Go 协程计算平方和，而在另一个协程计算立方和，最后在 Go 主协程把平方和与立方和相加。
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	//这两个函数分别在单独的协程里运行，每个函数都有传递信道的参数，以便写入数据
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	//创建了一个信道 ch，接着在下一行 ch <- 5，把 5 发送到这个信道。对于本程序，没有其他的协程从 ch 接收数据。于是程序触发 panic，
	//死锁 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。如果没有的话，程序就会在运行时触发 panic，形成死锁。
	//ch := make(chan int)
	//ch <- 5

}
