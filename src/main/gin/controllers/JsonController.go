package controllers

import (
	"github.com/gin-gonic/gin"
	"main/gin/models"
	"net/http"
)

/*
响应json数据
*/
func Json(ginServer *gin.Engine) {

	//1.返回普通数据类型
	ginServer.GET("/json1", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "success")
	})

	//2.map
	ginServer.GET("json2", func(context *gin.Context) {
		//2.1 使用gin自身的gin.h
		context.JSONP(http.StatusOK, gin.H{
			"code":    200,
			"message": "成功"})

		//2.2 自定义map
		//创建了一个变量  m ，它是一个空的映射，其中键的类型为  string ，值的类型为  any 。
		m := map[string]any{} // any是 Go 语言中的一种空接口类型，可以存储任意类型的值。
		m["code"] = 200
		m["message"] = "success"
		context.JSONP(http.StatusOK, m)

	})

	//3. 结构体
	ginServer.GET("json3", func(context *gin.Context) {
		//user结构体
		user := models.User{"001", "Alice", "6666"}
		context.JSONP(http.StatusOK, user)
	})

	//4.切片结构体
	ginServer.GET("json4", func(context *gin.Context) {
		user1 := models.User{"003", "Bob", "9898"}
		user2 := models.User{"004", "Eric", "6556"}
		//切片结构体
		users := make([]models.User, 2)
		users[0] = user1
		users[1] = user2

		context.JSONP(http.StatusOK, users)
	})
}
