package config

import "MinimalistTiktok/dao"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	dao.UserInfo `json:"user"`
}

type VideoListResponse struct {
	Response
	Video []dao.Video `json:"video_list"`
}

type CommentResponse struct {
	Response
	Comment dao.Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []dao.Comment `json:"comment_list"`
}
