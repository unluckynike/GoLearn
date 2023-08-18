package main

//$ curl -i http://localhost:9999/
//HTTP/1.1 200 OK
//Date: Mon, 12 Aug 2019 16:52:52 GMT
//Content-Length: 18
//Content-Type: text/html; charset=utf-8
//<h1>Hello Gee</h1>
//
//$ curl "http://localhost:9999/hello?name=geektutu"
//hello geektutu, you're at /hello
//
//$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
//{"password":"1234","username":"geektutu"}
//
//$ curl "http://localhost:9999/xxx"
//404 NOT FOUND: /xxx

import (
	"net/http"

	"gee"
)

/*
gee.go简单了不少。最重要的还是通过实现了 ServeHTTP 接口，接管了所有的 HTTP 请求。
*/
func main() {

	//创建实例
	r := gee.New()

	//传内嵌HTML
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee 这是 Contex的 HTML方法</h1>")
	})

	//传字符串
	r.GET("/hello", func(ctx *gee.Context) {
		// http://localhost:9999/hello?name=Lucis
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	//传json
	//localhost:9999/login?username=geektutu&password=1234
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.PUT("/add", func(ctx *gee.Context) {
		data := []byte{'h', 'e', 'l', 'l', 'o'}
		ctx.Data(http.StatusOK, data)
	})

	//启动
	r.Run(":9999")
}
