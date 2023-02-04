package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	service.FavoriteAction(c)
}
