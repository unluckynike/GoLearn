package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/gin/controllers"
)

func main() {

	fmt.Println("Gin Start。。。。。。。。。。。。")

	//创建ginserver
	ginserver := gin.Default()
	//创建服务
	//ginserver.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello Gin")
	//
	//})

	//路由
	controllers.Index(ginserver)
	//restfull API风格
	controllers.User(ginserver)
	//Json数据响应
	controllers.Json(ginserver)
	//请求参数传递
	controllers.Parameter(ginserver)
	//重定向
	controllers.Redirect(ginserver)
	//404页面处理
	controllers.Error(ginserver)

	fmt.Println("成功启动 端口8000")

	//加载静态页面位置
	//同级目录
	//ginserver.LoadHTMLGlob("D://Project//GoCode//GoLearn//src//main//gin//templates/*")
	//可以找到子目录
	ginserver.LoadHTMLGlob("D://Project//GoCode//GoLearn//src//main//gin//templates/**/*")

	//ginserver.LoadHTMLGlob("templates/*")   ??

	//配置静态资源地址 告诉ginserver这是静态资源 直接下载即可 不用去查找动态路由部分
	ginserver.Static("/static", "./static")

	//启动服务
	ginserver.Run("127.0.0.1:8000")
}
