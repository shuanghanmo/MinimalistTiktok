package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	service.Register(c)
}

func Login(c *gin.Context) {
	service.Login(c)
}

func UserInfo(c *gin.Context) {
	service.GetUserInfo(c)
}
