package controller

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"MinimalistTiktok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	token := c.Query("token")

	info, _ := service.ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	userId := userInfo.ID

	actionType := c.Query("action_type")
	if actionType == "1" {
		videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		content := c.Query("comment_text")
		comment, err := service.AddComment(userId, videoId, content)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, config.Response{StatusCode: -1, StatusMsg: "添加评论失败"})
			return
		}
		c.JSON(http.StatusOK, config.CommentResponse{
			Response: config.Response{
				StatusCode: 0,
				StatusMsg:  "添加评论成功",
			},
			Comment: *comment,
		})
	} else {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		err := service.DeleteComment(commentId, userId)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, config.Response{StatusCode: -1, StatusMsg: "删除评论失败"})
			return
		}
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		})
	}
}

func CommentList(c *gin.Context) {
	service.CommentList(c)
}
