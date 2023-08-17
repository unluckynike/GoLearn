package main

import (
	"fmt"
	"sync"
	"time"
)

/*
WaitGroup
WaitGroup 用于等待一批 Go 协程执行结束。程序控制会一直阻塞，直到这些协程全部执行完毕。
假设我们有 3 个并发执行的 Go 协程（由 Go 主协程生成）。
Go 主协程需要等待这 3 个协程执行结束后，才会终止。这就可以用 WaitGroup 来实现。
*/

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	//WaitGroup 是一个结构体类型，我们在第 18 行创建了 WaitGroup 类型的变量，其初始值为零值。WaitGroup 使用计数器来工作。
	//当我们调用 WaitGroup 的 Add 并传递一个 int 时，WaitGroup 的计数器会加上 Add 的传参。
	//要减少计数器，可以调用 WaitGroup 的 Done() 方法。Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
