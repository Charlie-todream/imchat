package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoadFrindRequest struct {
	Userid int64 `form:"userid,omitempty" json:"userid,omitempty" valid:"userid"`
}

type AddFrindRequest struct {
	Userid int64 `form:"userid,omitempty" json:"userid,omitempty" valid:"userid"`
	Dstid  int64 `form:"dstid,omitempty" json:"dstid,omitempty" valid:"dstid"`
}

type CreateCommunityRequest struct {
	Cate    int64  `form:"cate,omitempty" json:"cate,omitempty" valid:"cate"`
	Ownerid int64  `form:"ownerid,omitempty" json:"ownerid,omitempty" valid:"ownerid"`
	Name    string `form:"name,omitempty" json:"name,omitempty" valid:"name"`
	Memo    string `form:"memo,omitempty" json:"memo,omitempty" valid:"memo"`
	Icon    string `form:"icon,omitempty" json:"icon,omitempty" valid:"icon"`
}

// 验证创建社群
func CreateCommunity(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":    []string{"required"},
		"ownerid": []string{"required"},
	}
	messages := govalidator.MapData{
		"ownerid": []string{
			"required:Name为必填项，参数名称 name",
		},
		"name": []string{
			"required:name",
		},
	}

	err := validate(data, rules, messages)
	return err
}
