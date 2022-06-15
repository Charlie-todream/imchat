package main

import (
	"github.com/gin-gonic/gin"

	bootstrap "tobepower/chat/boostrap"
	bootconfig "tobepower/chat/config"
	"tobepower/chat/pkg/config"
)

func init() {
	// 加载config 目录下的配置信息
	bootconfig.Initialize()
}
func main() {

	config.InitConfig("env")
	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDB()
	// 路由
	bootstrap.SetupRoute(router)
	router.Run(":8080")
}
