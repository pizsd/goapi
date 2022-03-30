package requests

import (
	"github.com/gin-gonic/gin"
)
import "github.com/thedevsaddam/govalidator"

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	msg := govalidator.MapData{
		"email": []string{
			"required:邮箱是必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      msg,
	}

	return govalidator.New(opts).ValidateStruct()
}
