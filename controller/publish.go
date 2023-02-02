package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	service.Response
	VideoList []service.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, service.Response{
		StatusCode: 0,
		StatusMsg:  "该功能未完成",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
