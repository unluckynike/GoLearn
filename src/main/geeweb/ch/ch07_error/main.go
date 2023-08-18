package main

import (
	"net/http"

	"gee"
)

// 错误处理也可以作为一个中间件，增强web框架的能力。
func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello \n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"hailinlin"}
		c.String(http.StatusOK, names[100])
	})
	//在上面的代码中，我们为 gee 注册了路由 /panic，而这个路由的处理函数内部存在数组越界 names[100]，
	//如果访问 localhost:9999/panic，Web 服务就会宕掉。

	r.Run(":9999")
}
