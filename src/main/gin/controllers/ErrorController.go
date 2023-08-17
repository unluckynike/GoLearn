package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//错误页面处理

func Error(engine *gin.Engine) {
	engine.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusOK, "404.html", gin.H{"message": "页面找不见", "code": "404"})
	})
}
