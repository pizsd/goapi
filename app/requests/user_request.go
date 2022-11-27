package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/app/requests/validators"
	"goapi/pkg/auth"
	"strconv"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city,omitempty"`
	Introduction string `valid:"introduction" json:"introduction,omitempty"`
	Avatar       string `valid:"avatar" json:"avatar,omitempty"`
}

type UserUpdateEmailRequest struct {
	Email      string `valid:"email" json:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type UserUpdatePhoneRequest struct {
	Phone      string `valid:"phone" json:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

type UserUpdatePasswordRequest struct {
	Password           string `valid:"password" json:"password,omitempty"`
	NewPassword        string `valid:"new_password" json:"new_password,omitempty"`
	NewPasswordConfirm string `valid:"new_password_confirm" json:"new_password_confirm,omitempty"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {
	uid := strconv.FormatInt(auth.Uid(c), 10)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_dash", "between:5,20", "not_exists:users,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:15", "max_cn:240"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_dash:用户名只能暴行字母数字字符以及破折号和下划线",
			"between:用户名长度为 5-20 个字",
			"not_exists:名称已存在",
		},
		"city": []string{
			"min_cn:城市至少需要 2 个字",
			"max_cn:城市不能超过 20 个字",
		},
		"introduction": []string{
			"min_cn:描述长度需至少 15 个字",
			"max_cn:描述长度不能超过 240 个字",
		},
	}
	return validate(data, rules, messages)
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {
	userModel := auth.User(c)
	uid := strconv.FormatInt(userModel.ID, 10)

	rules := govalidator.MapData{
		"email":       []string{"required", "min:6", "max:30", "email", "not_exists:users,email," + uid, "not_in:" + userModel.Email},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email为必填项",
			"min:Email 至少需要 6 个字符",
			"max:Email 长不能超过 30 个字符",
			"email:Email 格式无效",
			"not_exists: Email 已被占用",
			"not_in: Email 不能与原 Email 一致",
		},
		"verify_code": []string{
			"required:验证码不能为空",
			"digits:手机号长度必须为 6 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {
	userModel := auth.User(c)
	uid := strconv.FormatInt(userModel.ID, 10)

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11", "not_exists:users,phone," + uid, "not_in:" + userModel.Phone},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号格式不正确",
			"not_exists: 手机号已被占用",
			"not_in: 手机号不能与原手机号一致",
		},
		"verify_code": []string{
			"required:验证码不能为空",
			"digits:手机号长度必须为 6 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:8"},
		"new_password":         []string{"required", "min:8"},
		"new_password_confirm": []string{"required", "min:8"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:旧密码长度需大于 8 个字符",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:新密码长度需大于 8 个字符",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 8 个字符",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.VerifyNewPasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)
	return errs
}
