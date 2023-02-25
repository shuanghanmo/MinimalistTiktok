package dao

import (
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:id" json:"id,omitempty"`
	UserId     int64     `gorm:"column:user_id" json:"-"`
	VideoId    int64     `gorm:"column:video_id" json:"-"`
	User       UserInfo  `gorm:"-" json:"user"`
	Content    string    `gorm:"column:content" json:"content,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "tb_comment"
}

func AddComment(userId int64, videoId int64, commentText string) (Comment, error) {
	var userInfo UserInfo
	DB.First(&userInfo, userId)
	time := time.Now()

	comment := Comment{
		UserId:     userId,
		VideoId:    videoId,
		User:       userInfo,
		Content:    commentText,
		CreateDate: time,
	}
	err := DB.Create(&comment).Error

	return comment, err
}

func DeleteComment(commentId int64) error {
	var comment Comment
	err := DB.Where("id = ?", commentId).Delete(&comment).Error
	return err
}

func QueryCommentList(userId int64, videoId int64) []Comment {
	var commentList []Comment
	var flag bool

	result := DB.Select("id", "user_id", "content", "create_date").Where("video_id = ?", videoId).Find(&commentList)

	n := result.RowsAffected
	if n == 0 {
		return nil
	}

	var i int64
	for i = 0; i < n; i++ {
		var userInfo UserInfo
		DB.Where("id = ?", commentList[i].UserId).Find(&userInfo)
		flag = IsFollow(userId, commentList[i].UserId)
		userInfo.IsFollow = flag
		commentList[i].User = userInfo
	}

	return commentList
}
