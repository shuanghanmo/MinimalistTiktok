package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList dao.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "该功能未完成",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		//VideoList: DemoVideos,
	})
}
