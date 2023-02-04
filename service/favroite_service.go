package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64) //string->int64
	//video := dao.QueryVideoByUserId(videoId)                    //string->int64
	action_type := c.Query("action_type")
	var msg string
	var err error

	//if action_type == "1" {
	//	err = dao.PlusOneFavorByUserIdAndVideoId(video.UserId, videoId)
	//} else {
	//	err = dao.MinusOneFavorByUserIdAndVideoId(video.UserId, videoId)
	//}
	if err == nil {
		msg = "操作成功"
	} else {
		msg = "操作失败"
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  msg,
	})
}
