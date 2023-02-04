package middleware

import (
	"MinimalistTiktok/service"
	"MinimalistTiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var LIST = [...]string{"/douyin/user/login", "/douyin/user/register", "/douyin/feed/", "/douyin/publish/list/"}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		for i := 0; i < len(LIST); i++ {
			if strings.HasPrefix(url, LIST[i]) {
				c.Next()
				return
			}
		}
		token := c.Query("token")
		id := c.Query("id")
		_, claim, err := utils.ParseToken(token)
		// 检测token是否过期
		if err != nil {
			handleError("认证失败！token已过期", c)
			c.Abort()
			return
		}
		// 检测是否携带id，如果携带与claim里面的id比较
		if len(id) != 0 {
			if claim.Id != id {
				handleError("认证失败！你没有权限执行此操作", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func handleError(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, service.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}
