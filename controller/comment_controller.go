package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	//TODO 评论操作
}

func CommentList(c *gin.Context) {
	service.CommentList(c)
}
