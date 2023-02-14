package dao

import "time"

type Comment struct {
	Id         int64     `gorm:"column:id" json:"id,omitempty"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	VideoId    int64     `gorm:"column:user_id" json:"video_id"`
	Content    string    `gorm:"column:content" json:"content,omitempty"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date,omitempty"`
}

type CommentList struct {
	Id         int64     `json:"id,omitempty"`
	User       UserInfo  `json:"userInfo"`
	Content    string    `json:"content,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "tb_comment"
}

func QueryCommentListByVideoId(videoId int64) []CommentList {
	var comments []Comment
	var userInfo UserInfo

	result := DB.Select("id", "user_id", "content", "create_date").Where("video_id = ?", videoId).Find(&comments)

	n := result.RowsAffected
	if n == 0 {
		return nil
	}
	commentList := make([]CommentList, n)

	var i int64
	for i = 0; i < n; i++ {
		commentList[i].Id = comments[i].Id
		DB.First(&userInfo, comments[i].UserId)
		commentList[i].User = userInfo
		commentList[i].Content = comments[i].Content
		commentList[i].CreateDate = comments[i].CreateDate
	}

	return commentList
}
