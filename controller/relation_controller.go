package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	service.RelationAction(c)
}
