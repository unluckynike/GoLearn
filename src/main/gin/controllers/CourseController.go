package controllers

import (
	"github.com/gin-gonic/gin"
	"main/gin/filter"
	"net/http"
)

/*
路由组
在业务架构，我们通常会把一类相同类型的业务采用相同的前缀名字以便和其它类型的业务区分开，方便开发人员维护和设计。
例如，有关用户的一类业务，接口都是/user/* ，订单的一类业务，接口为/order/*
*/

func Course(engine *gin.Engine) {
	courseGroup := engine.Group("/course", filter.MyHandle()) //filter.MyHandle() 局部注册 在路由组中加入中间件
	//在gin框架中，我们可以把这一统一的前缀提取出来，定义为一个路由组，方便区分，代码也会更简介
	//访问的真实接口是
	//course/list
	//course/update
	//course/add
	//course/delete

	{
		courseGroup.GET("/list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "list"})
		})
		courseGroup.GET("/add", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "add"})
		})
		courseGroup.GET("/del", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "del"})
		})
		courseGroup.GET("/update", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "update"})
		})
	}

	//在某一个请求上加入中间件，中间件作为参数定义在请求接口后，真正的请求之前
	engine.GET("/filter", filter.MyHandle(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET",
		})
	})

	/*
		gin默认中间件
		gin.Default()默认使用了Logger和Recovery中间件
		Logger中间件将日志写入gin.DefaultWriter，即配置了GIN_MODE = release。
		Recovery中间件会recover任何panic。如果有panic，会写入500响应码。
		如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个gin服务。
	*/

}
