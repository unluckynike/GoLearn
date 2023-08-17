package main

import (
	"fmt"
	"time"
)

/*
缓冲信道

通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
无缓冲信道的容量默认为 0
ch := make(chan type, capacity)
*/

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("成功写入", i, "进入ch信道")
	}
	close(ch)
}

func main() {
	ch1 := make(chan string, 2)
	ch1 <- "naveen"
	ch1 <- "paul"
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	//长度 容量
	ch2 := make(chan string, 3)
	ch2 <- "AAA"
	ch2 <- "AAB"
	ch2 <- "AAC"
	fmt.Println("capacity is ", cap(ch2))
	fmt.Println("length is ", len(ch2))
	fmt.Println("read value ", <-ch2)
	fmt.Println("new length is ", len(ch2))

	//一个并发的 Go 协程来向信道写入数据，而 Go 主协程负责读取数据
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	//在这期间，write 协程在并发地运行。write 协程有一个 for 循环，依次向信道 ch 写入 0～4。
	//而缓冲信道的容量为 2，因此 write 协程里立即会向 ch 写入 0 和 1，接下来发生阻塞，直到 ch 内的值被读取。
	for v := range ch {
		fmt.Println("从ch信道中读到值", v)
		time.Sleep(2 * time.Second)
	}

}
