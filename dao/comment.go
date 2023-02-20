package dao

import (
	"fmt"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:id" json:"id,omitempty"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	VideoId    int64     `gorm:"column:video_id" json:"video_id"`
	Content    string    `gorm:"column:content" json:"content,omitempty"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date,omitempty"`
}

type CommentList struct {
	Id         int64     `json:"id,omitempty"`
	User       UserInfos `json:"userInfo"`
	Content    string    `json:"content,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (c Comment) TableName() string {
	return "tb_comment"
}

type CommentDao struct {
}

var commentDao *CommentDao

func NewCommentDaoInstance() *CommentDao {
	return commentDao
}

func (*CommentDao) AddComment(comment *Comment) error {
	err := DB.Create(&comment).Error
	if err != nil {
		fmt.Println("insert video into db ERROR")
	}
	return err
}
func (*CommentDao) SelectCommentByID(commentId int64) (*Comment, error) {
	var comment Comment
	err := DB.Table("tb_comment").Where("id = ?", commentId).First(&comment).Error
	if err != nil {
		fmt.Println("评论id不存在" + err.Error())
	}
	return &comment, err
}

func (*CommentDao) DeleteComment(comment *Comment) error {
	err := DB.Delete(comment).Error
	if err != nil {
		fmt.Println("评论删除失败 " + err.Error())
	}
	return err
}

func QueryCommentList(userId int64, videoId int64) []CommentList {
	var comments []Comment
	var userInfos UserInfos
	var flag bool

	result := DB.Select("id", "user_id", "content", "create_date").Where("video_id = ?", videoId).Find(&comments)

	n := result.RowsAffected
	if n == 0 {
		return nil
	}
	commentList := make([]CommentList, n)

	var i int64
	for i = 0; i < n; i++ {
		var userInfo UserInfo
		commentList[i].Id = comments[i].Id
		DB.Where("id = ?", comments[i].UserId).Find(&userInfo)
		flag = IsFollow(userId, comments[i].UserId)
		userInfos = SaveUserInfos(userInfo, flag)
		commentList[i].User = userInfos
		commentList[i].Content = comments[i].Content
		commentList[i].CreateDate = comments[i].CreateDate
	}

	return commentList
}
