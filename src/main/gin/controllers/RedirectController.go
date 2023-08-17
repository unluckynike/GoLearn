package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect(engine *gin.Engine) {
	// 站外 重定向
	engine.GET("/test", func(context *gin.Context) {
		fmt.Println("重定向到百度")
		//context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
		context.Redirect(http.StatusFound, "https://www.baidu.com")

	})

	// 站内 重定向
	engine.GET("/test/index", func(context *gin.Context) {
		fmt.Println("重定向index")
		context.Redirect(http.StatusMovedPermanently, "/index")

	})
}
