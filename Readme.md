用go gin gorm开发的一个聊天系统 支持 文字和图片 整体架构类似php的laravel框架 
后续加入jwt,按go1.13 Wrap方式处理错误 

// 多模板渲染
go get github.com/gin-contrib/multitemplate 

// 配置
go get github.com/spf13/cast
go get github.com/spf13/viper

// 日志 lumberjack 是一套滚动日志的实现方案，帮助我们管理日志文件。
go get go.uber.org/zap
go get gopkg.in/natefinch/lumberjack.v2


