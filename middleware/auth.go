package middleware

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var LIST = [...]string{"/douyin/user/login", "/douyin/user/register"}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.RequestURI
		for i := 0; i < len(LIST); i++ {
			if strings.HasPrefix(url, LIST[i]) {
				c.Next()
				return
			}
		}
		token := c.Query("token")
		if len(token) == 0 {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: 1,
				StatusMsg:  "认证失败！",
			})
			return
		}
		cache, ok := service.ConcurrentMap.Load(token)
		if !ok || cache == nil {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: 1,
				StatusMsg:  "认证失败！",
			})
			return
		}
		c.Next()
	}
}
