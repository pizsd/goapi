package auth

import (
	"errors"
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
