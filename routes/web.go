package routes

import (
	"github.com/gin-gonic/gin"
	"tobepower/chat/handler"
)

// RgisterWebRoutes 注册网页相关路由

func RegisterWebRoutes(r *gin.Engine) {

	chatHandler := new(handler.Chat)
	r.GET("", chatHandler.Index)

	v1 := r.Group("/v1")

	authHandler := new(handler.Auth)
	auth := v1.Group("/auth")
	{
		auth.GET("/login", authHandler.Login)
		auth.POST("/login-do", authHandler.LoginStore)
		auth.GET("/register", authHandler.Register)
		auth.POST("/register-do", authHandler.RegisterStore)

	}

	contactHandler := new(handler.Contact)
	contact := v1.Group("/contact")
	{
		contact.POST("loadfriend", contactHandler.LoadFriend)
		contact.POST("loadcommunity", contactHandler.LoadCommunity)
		contact.POST("addfriend", contactHandler.Addfriend)
		contact.POST("joincommunity", contactHandler.JoinCommunity)
		contact.GET("createcom", contactHandler.Createcom)
		contact.POST("createcommunity", contactHandler.Createcommunity)
	}
	chatMessageHandler := new(handler.ChatMessage)
	chat := v1.Group("/chat")
	{
		chat.GET("message", chatMessageHandler.Chat)
	}

	attachHandler := new(handler.Attach)
	attach := v1.Group("/attach")
	{
		attach.POST("upload", attachHandler.UploadLocal)
	}

}
