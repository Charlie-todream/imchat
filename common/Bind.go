package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 绑定form 或json
func Bind(c *gin.Context, obj interface{}) bool {
	//contentType := c.Request.Header.Get("Content-Type")
	//if strings.ContainsAny(strings.ToLower(contentType),"application/json"){
	//	fmt.Println("json")
	//	return BindFromJson(c,obj)
	//
	//}
	//
	//if strings.Contains(strings.ToLower(contentType),"application/x-www-form-urlencoded"){
	//	fmt.Println("form")
	//	return   BindFromForm(c,obj)
	//}

	return AllType(c, obj)
}

func BindFromJson(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误,请确认请求格式是否正确",
			"error":   err.Error(),
		})
		return false
	}
	return true
}

func BindFromForm(c *gin.Context, obj interface{}) bool {
	if err := c.Bind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误,请确认请求格式是否正确",
			"error":   err.Error(),
		})
		return false
	}
	return true
}

func AllType(c *gin.Context, obj interface{}) bool {

	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误,请确认请求格式是否正确",
			"error":   err.Error(),
		})

		return false
	}
	return true
}
