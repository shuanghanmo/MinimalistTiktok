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
)

type VideoListResponse struct {
	Response
	VideoList []dao.VideoList `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	_, claim, err := utils.ParseToken(token)
	userId := claim.Uid
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, Response{
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
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	coverPath, videoPath := utils.GetCover(finalName)

	// 存数据库
	video := dao.Video{
		Id:       videoId,
		UserId:   userId,
		PlayUrl:  videoPath,
		CoverUrl: coverPath,
		Title:    title,
		//IsFavorite:    false,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err = dao.NewVideoDaoInstance().AddVideo(video)
	if err == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  filename + " uploaded successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 1,
			StatusMsg:  filename + " uploaded Failed",
		})
	}

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
