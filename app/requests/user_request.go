package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goapi/pkg/auth"
	"strconv"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city,omitempty"`
	Introduction string `valid:"introduction" json:"introduction,omitempty"`
	Avatar       string `valid:"avatar" json:"avatar,omitempty"`
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
