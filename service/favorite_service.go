package service

import (
	"MinimalistTiktok/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	info, _ := ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	userId := userInfo.ID
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "信息格式转化错误！",
		})
		return
	}
	//flag := dao.IsFavorVideo(videoId, userId)
	actionTypestr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypestr, 10, 64)
	if actionType == 1 {
		err := dao.PlusOneFavorByUserIdAndVideoId(userId, videoId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "发生异常错误，请稍后访问！",
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功！",
		})
	} else {
		err := dao.MinusOneFavorByUserIdAndVideoId(userId, videoId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "发生异常错误，请稍后访问！",
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功！",
		})
	}
}

func FavoriteList(c *gin.Context) {

}
