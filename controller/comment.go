package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	service.Response
	CommentList []service.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	service.Response
	Comment service.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    service.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
