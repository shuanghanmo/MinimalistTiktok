package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	service.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	service.Response
	User dao.UserInfo `json:"user_info"`
}

func Register(c *gin.Context) {
	service.Register(c)
}

func Login(c *gin.Context) {
	service.Login(c)
}

func UserInfo(c *gin.Context) {
	service.GetUserInfo(c)
}
