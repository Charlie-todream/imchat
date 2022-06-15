package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Chat struct {
}

func (ctrl *Chat) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "chat/index.html", gin.H{
		"title": "index",
	})
}
