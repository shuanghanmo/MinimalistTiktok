package config

import "MinimalistTiktok/dao"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []dao.VideoList `json:"video_list"`
}

type CommentListResponse struct {
	Response
	CommentList []dao.CommentList `json:"comment_list"`
}
