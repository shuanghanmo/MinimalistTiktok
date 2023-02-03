package controller

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	service.Publish(c)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	service.PublishList(c)
}
