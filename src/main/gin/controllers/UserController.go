package controllers

import "github.com/gin-gonic/gin"

func User(ginServer *gin.Engine) {

	/*
		restful api
	*/

	//get
	ginServer.GET("/user", func(context *gin.Context) {
		//传递的是json数据
		context.JSONP(200, gin.H{"message": "get"})
	})
	//post
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSONP(200, gin.H{"message": "post"})
	})
	//del
	ginServer.DELETE("/user", func(context *gin.Context) {
		context.JSONP(200, gin.H{"message": "delete"})
	})
	//put
	ginServer.PUT("/user", func(context *gin.Context) {
		context.JSONP(200, gin.H{"message": "put"})
	})
}
