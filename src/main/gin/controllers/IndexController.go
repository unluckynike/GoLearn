package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ginServer *gin.Engine) {

	ginServer.GET("/index", func(context *gin.Context) {
		//页面是index字符
		//context.String(200, "index")

		context.JSONP(200, gin.H{"message": "this is json string"})
	})

	//配置一个路由首页 相应数据
	ginServer.GET("/hello", func(context *gin.Context) {

		session := sessions.Default(context)
		user := session.Get("user")

		context.HTML(
			http.StatusOK, //200
			"index.html",
			gin.H{
				"message": "Gin hello",
				"code":    "200",
				"user":    user})
	})

}
