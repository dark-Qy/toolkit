// name:qy
// func:用户认证中间件

package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"toolkit/config"
)

// Claims 是一个结构体，继承jwt.StandardClaims
type Claims struct {
	Username string `json:"userName"`
	jwt.StandardClaims
}

// GenerateJWT 生成JWT令牌
func GenerateJWT(username string) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 对身份信息进行加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtSecret)

	return tokenString, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization字段
		authHeader := c.GetHeader("Authorization")
		// 如果没有Authorization字段，则返回错误
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// 解析Authorization字段，获取token字符串
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &Claims{}

		// 使用配置中的JWT密钥解析token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		// 如果解析错误或者token无效，继续返回错误
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
