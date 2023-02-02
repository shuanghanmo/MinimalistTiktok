package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var tempChat = map[string][]service.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	service.Response
	MessageList []service.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
