// name:qy
// func:用户数据模型

package models

import (
	"errors"
	"toolkit/dao"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser 新建用户
func CreateUser(user *User) (err error) {
	err = dao.DB.Create(&user).Error
	return nil
}

func GetUser(user *User) (err error) {
	err = dao.DB.Debug().Where("username=?", user.Username).First(&user).Error
	if user.Username == "" {
		return errors.New("user not found")
	}
	return nil
}
