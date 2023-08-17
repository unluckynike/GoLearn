package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
获取请求参数
*/
func Parameter(ginServer *gin.Engine) {

	//1.ge 编写请求 基于key-value
	//http://127.0.0.1:8000/keyvalue/info?userid=666&username=Lrise
	ginServer.GET("/keyvalue/info", func(context *gin.Context) {
		//1 获取传统的基于key-value  Query 查询方法获取
		userid := context.Query("userid")
		username := context.Query("username")

		context.JSONP(http.StatusOK, gin.H{"userid": userid, "username": username})
	})

	//2.接受restful风格 的参数
	//127.0.0.1:8000/params/info/666/Linlin
	//相对于seo会更友好
	ginServer.GET("/params/info/:userid/:username", func(context *gin.Context) {
		//使用Param
		userid := context.Param("userid")
		username := context.Param("username")

		context.JSONP(http.StatusOK, gin.H{"userid": userid, "username": username})
	})

	//3.post请求 -- form表单
	//127.0.0.1:8000/params/save/
	ginServer.POST("/params/save", func(context *gin.Context) {
		// form 表单的方式 使用PostForm
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.JSONP(http.StatusOK, gin.H{"username": username, "password": password})
	})

	//4.post 请求 json参数的方式
	//127.0.0.1:8000/params/json  row中设置json 可以有多个键值对
	ginServer.POST("/params/json", func(context *gin.Context) {
		//GetRawData 原理 从request.Body中获取数据，返回的是一个[]byte
		data, _ := context.GetRawData()
		//开始定义map或者结构体
		var m map[string]any
		//把前台传递过来的json字符串转换map
		_ = json.Unmarshal(data, &m)

		context.JSONP(http.StatusOK, m)
	})

}
