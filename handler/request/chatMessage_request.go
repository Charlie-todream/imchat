package request

type ChatQuery struct {
	Id    int64  `form:"id"`
	Token string `form:"token"`
}
