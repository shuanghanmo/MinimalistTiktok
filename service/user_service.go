package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	"time"
)

var ConcurrentMap = sync.Map{}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	dao.UserInfo `json:"user"`
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	password = utils.Md5Crypt(password, "douyin_simple")
	user := dao.QueryByUser(username, password)
	if user.ID == 0 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "帐号或密码错误！",
		})
		return
	}
	userInfo := dao.QueryUserInfoById(user.ID)
	token, _ := utils.Award(user.ID)
	put(token, userInfo)
	// 保存用户对应的信息
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "登录成功！"},
		UserId:   user.ID,
		Token:    token,
	})
}

func put(key string, userInfo *dao.UserInfo) {
	var userCache = dao.UserInfo{
		userInfo.ID,
		userInfo.Name,
		userInfo.FollowCount,
		userInfo.FollowerCount,
		userInfo.IsFollow,
	}
	ConcurrentMap.Store(key, userCache)
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := dao.QueryByUserName(username)
	var userInfo *dao.UserInfo
	if user.ID != 0 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "该用户名已注册！",
		})
		return
	}
	password = utils.Md5Crypt(password, "douyin_simple")
	// 添加事务
	err := dao.DB.Transaction(func(tx *gorm.DB) error {
		user = &dao.User{UserName: username, PassWord: password}
		if err := tx.Model(&dao.User{}).Create(user).Error; err != nil {
			tx.Rollback()
			return err
		}
		name := "douyin_simple_" + time.Now().Format("2006:01:02") + "_" + username
		userInfo = &dao.UserInfo{ID: user.ID, Name: name}
		if err := tx.Model(&dao.UserInfo{}).Create(userInfo).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "发生异常错误！",
		})
		return
	}
	token, _ := utils.Award(user.ID)
	// 保存用户对应的信息
	put(token, userInfo)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "注册成功！"},
		UserId:   user.ID,
		Token:    token,
	})
}

func GetUserInfo(c *gin.Context) {
	token := c.Query("token")
	info, _ := ConcurrentMap.Load(token)
	userInfo := info.(dao.UserInfo)
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Response{StatusCode: 0, StatusMsg: "获取信息成功！"},
		UserInfo: userInfo,
	})
}
