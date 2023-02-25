package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	service.CommentAction(c)
}

func CommentList(c *gin.Context) {
	service.CommentList(c)
}
