package main

import (
	"fmt"
	"time"
)

/*
Go 协程
*/

func hello() {
	fmt.Println("Hello world goroutine")
}

/*
启动多个协程
*/
func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func printAAA() {
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("AAA ")
	}
}
func main() {

	/*
		Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。与线程相比，创建一个 Go 协程的成本很小。因此在 Go 应用中，常常会看到有数以千计的 Go 协程并发地运行。

		Go 协程相比于线程的优势
		相比线程而言，Go 协程的成本极低。堆栈大小只有若干 kb，并且可以根据应用的需求进行增减。而线程必须指定堆栈的大小，其堆栈是固定不变的。
		Go 协程会复用（Multiplex）数量更少的 OS 线程。即使程序有数以千计的 Go 协程，也可能只有一个线程。如果该线程中的某一 Go 协程发生了阻塞（比如说等待用户输入），那么系统会再创建一个 OS 线程，并把其余 Go 协程都移动到这个新的 OS 线程。所有这一切都在运行时进行，作为程序员，我们没有直接面临这些复杂的细节，而是有一个简洁的 API 来处理并发。
		Go 协程使用信道（Channel）来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件（Race Condition）。信道可以看作是 Go 协程之间通信的管道。
	*/

	/*
		两个 Go 协程的主要性质。
		启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
		如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	*/
	go hello()
	time.Sleep(1 * time.Second) //Go 主协程休眠了 1 秒。因此在主协程终止之前，调用 go hello() 就有足够的时间来执行了
	fmt.Println("main function")

	//启动多个协程
	go numbers()
	go alphabets()
	go printAAA()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
