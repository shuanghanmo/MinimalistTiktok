package service

import (
	"MinimalistTiktok/config"
	"MinimalistTiktok/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	info, _ := ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	userId := userInfo.ID
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
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
		err := dao.FollowAction(userId, toUserId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, config.Response{
				StatusCode: 1,
				StatusMsg:  "发生异常错误，请稍后访问！",
			})
			return
		}
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		})
	} else {
		err := dao.UnFollowAction(userId, toUserId)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, config.Response{
				StatusCode: 1,
				StatusMsg:  "发生异常错误，请稍后访问！",
			})
			return
		}
		c.JSON(http.StatusOK, config.Response{
			StatusCode: 0,
			StatusMsg:  "取消关注成功！",
		})
	}
}
