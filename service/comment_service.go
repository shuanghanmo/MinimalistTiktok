package service

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	//TODO 评论操作
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	info, _ := ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	userId := userInfo.ID

	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 1,
			StatusMsg:  "信息格式转化错误！",
		})
		return
	}

	var commentList = dao.QueryCommentList(userId, videoId)

	c.JSON(http.StatusOK, config.CommentListResponse{
		Response: config.Response{
			StatusCode: 0,
			StatusMsg:  "评论列表已刷新",
		},
		CommentList: commentList,
	})
}
