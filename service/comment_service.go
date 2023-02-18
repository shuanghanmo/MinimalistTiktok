package service

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"MinimalistTiktok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func AddComment(userId int64, videoId int64, content string) (*config.Comment, error) {
	user, err := dao.QueryById(userId)
	if err != nil || user.ID == 0 {
		return nil, errors.New("userId查询失败")
	}
	video, err := dao.NewVideoDaoInstance().SelectById(videoId)
	if err != nil || video.Id == 0 {
		return nil, errors.New("videoId查询失败")
	}

	comment := &dao.Comment{
		Id:         utils.GenSnowflake(),
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateDate: time.Now(),
		//CreateDate: time.Now().Month().String() + "-" + strconv.Itoa(time.Now().Day()),
	}
	err = dao.NewCommentDaoInstance().AddComment(comment)
	if err != nil {
		return nil, err
	}
	dao.NewVideoDaoInstance().IncrCommentCount(videoId)
	voComment, err := fillComment(user, comment, userId)
	if err != nil {
		return nil, err
	}
	return voComment, nil
}

func DeleteComment(commentId int64, userId int64) error {
	var err error
	comment, err := dao.NewCommentDaoInstance().SelectCommentByID(commentId)
	if err != nil || comment.Id == 0 {
		return errors.New("查找数据库失败，comment或许不存在")
	}
	if comment.UserId != userId {
		return errors.New("没有权限删除")
	}
	err = dao.NewCommentDaoInstance().DeleteComment(comment)
	if err != nil {
		return err
	}
	dao.NewVideoDaoInstance().DecrCommentCount(comment.VideoId)
	return nil
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

func fillComment(user *dao.User, comment *dao.Comment, loginUserId int64) (*config.Comment, error) {
	isFollow := dao.IsFollow(loginUserId, user.ID)
	userinfo := dao.QueryUserInfoById(user.ID)

	voUser := config.User{
		Id:            userinfo.ID,
		Name:          userinfo.Name,
		FollowCount:   userinfo.FollowCount,
		FollowerCount: userinfo.FollowerCount,
		IsFollow:      isFollow,
	}

	return &config.Comment{
		Id:         comment.Id,
		User:       voUser,
		Content:    comment.Content,
		CreateDate: comment.CreateDate.Month().String() + "-" + strconv.Itoa(time.Now().Day()),
	}, nil
}
