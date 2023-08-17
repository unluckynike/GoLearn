package main

import "fmt"

/*
单向信道
*/

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	//创建了唯送（Send Only）信道 sendch。chan<- int 定义了唯送信道，因为箭头指向了 chan
	sendch := make(chan<- int)
	go sendData(sendch)
	//fmt.Println(<-sendch)

	//创建了一个双向信道 cha1
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)

	//关闭信道和
	ch := make(chan int)
	go producer(ch)
	//producer 协程会从 0 到 9 写入信道 chn1，然后关闭该信道
	//for {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("Received ", v, ok)
	//}
	//使用 for range 遍历信道
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
