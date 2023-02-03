package middleware

import (
	"MinimalistTiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func FormatCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if len(username) == 0 || len(password) == 0 || len(password) > 32 {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: 1,
				StatusMsg:  "帐号或密码格式错误！",
			})
			c.Abort()
			return
		}
		ok, _ := regexp.MatchString(`^([\w._\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, username)
		if !ok {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: 1,
				StatusMsg:  "邮箱格式错误！",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
