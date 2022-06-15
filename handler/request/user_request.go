package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Mobile   string `form:"mobile,omitempty" json:"mobile,omitempty" valid:"mobile"`
	Password string `form:"password,omitempty" json:"password,omitempty" valid:"password"`
}

// LoginByPhone 验证表单，返回长度等于零即通过
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile":   []string{"required"},
		"password": []string{"required"},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 mobile",
			//"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码必填",
		},
	}

	err := validate(data, rules, messages)
	return err
}

type RegisterRequest struct {
	Mobile   string `form:"mobile,omitempty" json:"mobile,omitempty" valid:"mobile"`
	Password string `form:"password,omitempty" json:"password,omitempty" valid:"password"`
	Nickname string `form:"nickname,omitempty" json:"nickname,omitempty" valid:"nickname"`
	Avatar   string `form:"avatar,omitempty" json:"avatar,omitempty" valid:"avatar"`
	Sex      string `form:"sex,omitempty" json:"sex,omitempty" valid:"sex"`
	Salt     string `form:"salt,omitempty" json:"salt,omitempty" valid:"salt"`
	Online   int    `form:"online,omitempty" json:"online,omitempty" valid:"online"`
	Token    string `form:"token,omitempty" json:"token,omitempty" valid:"token"`
	Memo     string `form:"memo,omitempty" json:"memo,omitempty" valid:"memo"`
}

func ReigsterUser(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile":   []string{"required"},
		"password": []string{"required"},
		"sex":      []string{"required"},
	}

	messages := govalidator.MapData{

		"mobile": []string{
			"required:手机号为必填项，参数名称 mobile",
		},

		"password": []string{
			"required:密码必填",
		},
	}

	err := validate(data, rules, messages)
	return err

}
