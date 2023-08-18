package main

import (
	"fmt"
	"log"
	"net/http"
)

// Go语言内置了 net/http库，封装了HTTP网络编程的基础的接口
func main() {
	//http://localhost:9999/
	http.HandleFunc("/", indexHandler)
	//http://localhost:9999/hello    响应是请求头(header)中的键值对信息。
	http.HandleFunc("/hello", helloHandler)

	//nil 代表使用标准库中的实例处理
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
