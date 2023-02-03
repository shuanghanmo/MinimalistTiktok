package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	service.Response
	VideoList dao.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	service.Publish(c)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	service.PublishList(c)
}
