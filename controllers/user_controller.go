// name:qy
// func:用户控制器文件

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toolkit/middlewares"
	"toolkit/models"
)

func UserRegisterHandler(c *gin.Context) {
	// 根据Json信息绑定结构体
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	fmt.Println("user is :", user)

	// 创建对应用户信息
	err = models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1001, "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 0, "userName": user.Username})
	}
}

func UserLoginHandler(c *gin.Context) {
	// 根据Json信息绑定结构体
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	// 查询对应用户信息
	err = models.GetUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1001, "error": err.Error()})
	}

	// 用户认证成功，生成JWT
	token, err := middlewares.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 1002,
			"error":  "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"userName": user.Username,
		"token":    token,
	})
}
