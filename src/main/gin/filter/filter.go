package filter

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
gin框架允许开发者在处理请求时，加入用户自己设计的函数。
它用于处理一些公共的业务逻辑。例如登录认证、权限校验、数据分页、记录日志、耗时统计等。
也就是我们所说的中间件。
必须是一个gin.HanderFunc类型。
*/

func MyHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		//可以通过两个方法来决定是继续Netx还是Abort

		fmt.Println("middle start")

		context.Next()
		//fmt.Println("请求 ----- start")
		//if true {
		//	// 请求继续向下执行
		//	ctx.Next()
		//} else {
		//	// 中断
		//	ctx.Abort()
		//}
		//fmt.Println("请求 ----- end")
		fmt.Println("filter ending")
	}
}

func RegFulter(engine *gin.Engine) {
	//注册中间件 全局注册
	//engine.Use(MyHandle())

}
