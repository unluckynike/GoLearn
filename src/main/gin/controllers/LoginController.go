package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main/gin/models"
	"net/http"
)

//session会话

func Login(engine *gin.Engine) {

	engine.POST("/login", func(context *gin.Context) {
		//获取登录用户信息
		username := context.PostForm("username")
		password := context.PostForm("password")
		//创建一个用户结构体
		user := models.User{"001", username, password}
		//把结构体放入session
		session := sessions.Default(context)
		session.Set("user", user)
		//返回
		context.JSONP(http.StatusOK, gin.H{"user": user, "status": 200})

	})
}
