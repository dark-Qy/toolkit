// name:qy
// func:路由

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toolkit/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册测试路由
	r.GET("/", func(c *gin.Context) {
		fmt.Println("test ok!")
	})
	// 注册用户相关的路由
	UserGroup := r.Group("user")
	{
		// 注册
		UserGroup.POST("/register", controllers.UserRegisterHandler)
		// 登录
		UserGroup.POST("/login", controllers.UserLoginHandler)
	}

	return r
}
