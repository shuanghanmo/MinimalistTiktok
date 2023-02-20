package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	service.RelationAction(c)
}

func FollowList(c *gin.Context) {
	service.FollowList(c)
}

func FollowerList(c *gin.Context) {
	service.FollowerList(c)
}
