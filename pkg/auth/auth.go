package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"goapi/app/models/user"
)

func Attempt(loginId, password string) (user.User, error) {
	userModel := user.GetByMulti(loginId)
	if userModel.ID == 0 {
		return user.User{}, errors.New("帐号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}
	return userModel, nil
}

func User(c *gin.Context) user.User {
	userModel := c.MustGet("user").(user.User)
	return userModel
}

func Uid(c *gin.Context) int64 {
	uidStr := c.GetString("uid")
	return cast.ToInt64(uidStr)
}
