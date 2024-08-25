// name:qy
// func:配置文件

package config

import (
	"os"
)

var (
	JwtSecret = []byte("my_secret_key")
)

func InitConfig() {
	// 读取JWT_SECRET作为secret，如果不存在则用默认值
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		JwtSecret = []byte(secret)
	}
}
