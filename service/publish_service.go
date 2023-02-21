package service

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"MinimalistTiktok/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	_, claim, err := utils.ParseToken(token)
	userId := claim.Uid
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")

	//获取视频的格式
	filename := filepath.Base(data.Filename)
	videoNameList := strings.Split(filename, ".")
	format := videoNameList[1]
	//使用videoId作为视频的名字
	videoId := utils.GenSnowflake()
	finalName := fmt.Sprintf("%d.%s", videoId, format)

	saveFile := filepath.Join(config.VideosImagePath, finalName)
	if err = c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	coverPath, videoPath := utils.GetCover(finalName)

	// 存数据库
	video := dao.Video{
		Id:            videoId,
		UserId:        userId,
		PlayUrl:       videoPath,
		CoverUrl:      coverPath,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err = dao.NewVideoDaoInstance().AddVideo(video)
	if err == nil {
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 0,
			StatusMsg:  filename + " uploaded successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, config.Response{
			StatusCode: 1,
			StatusMsg:  filename + " uploaded Failed",
		})
	}

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) //string->int64

	var video = dao.QueryPublishListByUserId(userId)

	c.JSON(http.StatusOK, config.VideoResponse{
		Response: config.Response{
			StatusCode: 0,
			StatusMsg:  "发布列表已刷新",
		},
		Video: video,
	})
}

func Feed(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) //string->int64

	rawTimestamp := c.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err != nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}

	var videoList = dao.QueryVideoListByLimitAndTime(userId, 30, latestTime)

	c.JSON(http.StatusOK, config.VideoResponse{
		Response: config.Response{
			StatusCode: 0,
			StatusMsg:  "视频列表已刷新",
		},
		Video: videoList,
	})
}
