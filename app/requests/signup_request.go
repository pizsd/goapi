package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/app/requests/validators"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

type SignupUsingPhoneRequest struct {
	Name            string `json:"name,omitempty" valid:"name"`
	Phone           string `json:"phone,omitempty" valid:"phone"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
}

type SignupUsingEmailRequest struct {
	Name            string `json:"name,omitempty" valid:"name"`
	Email           string `json:"email,omitempty" valid:"email"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	msg := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, msg)
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
	return validate(data, rules, msg)
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:5,20"},
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"password":         []string{"required", "between:8,20"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"name:用户名不能为空",
			"alpha_num:用户名必须由字母数字组成",
			"between:用户名必须由 5-20 位字母数字组成",
		},
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码不能为空",
			"between:密码必须由 8-20 位字符组成",
		},
		"password_confirm": []string{
			"required:确认密码不能为空",
		},
		"verify_code": []string{
			"required:验证码不能为空",
			"digits:手机号长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.VerifyPasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.VerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:5,20"},
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"password":         []string{"required", "between:8,20"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"name:用户名不能为空",
			"alpha_num:用户名必须由字母数字组成",
			"between:用户名必须由 5-20 位字母数字组成",
		},
		"email": []string{
			"required:邮箱是必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码不能为空",
			"between:密码必须由 8-20 位字符组成",
		},
		"password_confirm": []string{
			"required:确认密码不能为空",
		},
		"verify_code": []string{
			"required:验证码不能为空",
			"digits:手机号长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*SignupUsingEmailRequest)
	errs = validators.VerifyPasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.VerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}
