package service

import (
	"MinimalistTiktok/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []dao.VideoList `json:"video_list"`
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
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) //string->int64
	//token := c.Query("token")

	var videoList = dao.QueryPublishListByUserId(userId)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "发布列表已刷新",
		},
		VideoList: videoList,
	})
}
