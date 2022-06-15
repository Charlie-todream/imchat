package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	"tobepower/chat/common"
	"tobepower/chat/domain/models"
	"tobepower/chat/domain/repository"
	"tobepower/chat/handler/request"
	"tobepower/chat/pkg/database"
	"tobepower/chat/pkg/response"
)

type Auth struct {
}

func (ctrl *Auth) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", gin.H{
		"title": "Posts",
	})
}
func (ctrl *Auth) LoginStore(c *gin.Context) {
	requstForm := request.LoginByPhoneRequest{}
	if ok := request.Validate(c, &requstForm, request.LoginByPhone); !ok {
		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "表單验证失败",
		})
		return
	}

	// 登陆

	userRepository := repository.NewUserRepository(database.DB)

	user, err := userRepository.FindUser(requstForm.Mobile)

	if err != nil {
		response.Unauthorized(c, "登陆失败")
		return
	}
	salt := user.Salt
	password := user.Password
	requstForm.Password = common.MakePasswd(requstForm.Password, salt)
	if password == requstForm.Password {
		str := fmt.Sprintf("%d", time.Now().Unix())
		token := common.MD5Encode(str)
		// 更新token
		userRepository.UpdateToken(token, user.Id)
		response.JSON(c, gin.H{

			"code":  200,
			"msg":   "登陆成功",
			"token": token,
			"id":    user.Id,
		})
	} else {
		response.JSON(c, gin.H{

			"Code": 0,
			"Msg":  "登录失败用户名或密码错误",
		})
	}

}

func (ctr *Auth) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register.html", gin.H{
		"title": "Reister",
	})
}

func (ctrl *Auth) RegisterStore(c *gin.Context) {
	requstForm := request.RegisterRequest{}
	if ok := request.Validate(c, &requstForm, request.ReigsterUser); !ok {
		response.JSON(c, gin.H{

			"Code": 0,
			"Msg":  "表單验证失败",
		})
		return
	}

	userRepository := repository.NewUserRepository(database.DB)
	// 检查手机号是否存在
	user, _ := userRepository.FindUser(requstForm.Mobile)

	if user.Id > 0 {
		response.JSON(c, gin.H{
			"code": 0,
			"msg":  "用户已经注册",
		})
		return
	}
	requstForm.Salt = fmt.Sprintf("%08d", rand.Int31())
	requstForm.Password = common.MakePasswd(requstForm.Password, requstForm.Salt)

	insertUser := &models.User{}
	common.SwapTo(requstForm, insertUser)

	// 注册

	resId, err := userRepository.UserCreate(insertUser)

	if err != nil {
		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "注册失败",
		})
		return
	}

	if resId > 0 {
		response.JSON(c, gin.H{
			"code": 200,
			"msg":  "注册成功",
			"id":   resId,
			"data": requstForm,
		})
	} else {
		response.JSON(c, gin.H{

			"code": 0,
			"msg":  "注册失败",
		})
		return
	}

}
