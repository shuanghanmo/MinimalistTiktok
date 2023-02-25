package service

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	var msg string
	var comment dao.Comment

	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	info, _ := ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	userId := userInfo.ID
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 1,
			StatusMsg:  "信息格式转化错误！",
		})
		return
	}

	actionTypestr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypestr, 10, 64)
	if actionType == 1 {
		commentText := c.Query("comment_text")

		comment, err = dao.AddComment(userId, videoId, commentText)
		if err != nil {
			log.Println(err)
			msg = "发生异常错误，请稍后访问！"
		} else {
			msg = "评论成功！"
		}
	} else {
		commentIdstr := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentIdstr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, config.Response{
				StatusCode: 1,
				StatusMsg:  "信息格式转化错误！",
			})
			return
		}
		err = dao.DeleteComment(commentId)
		if err != nil {
			log.Println(err)
			msg = "发生异常错误，请稍后访问！"
		} else {
			msg = "删除评论成功！"
		}
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 0,
			StatusMsg:  msg,
		})
		return
	}
	c.JSON(http.StatusOK, config.CommentResponse{
		Response: config.Response{
			StatusCode: 0,
			StatusMsg:  msg,
		},
		Comment: comment,
	})
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
