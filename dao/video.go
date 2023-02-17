package dao

import (
	"fmt"
	_ "gorm.io/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	UserId        int64  `gorm:"column:user_id"`
	PlayUrl       string `gorm:"column:play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
	Title         string `gorm:"column:title"`
}

type VideoList struct {
	Id            int64     `gorm:"column:id" json:"id,omitempty"`
	Author        UserInfos `gorm:"column:author" json:"author"`
	PlayUrl       string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count" json:"comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite" json:"is_favorite"`
	Title         string    `gorm:"column:title" json:"title"`
}

type UserFavorVideo struct {
	UserId   int64 `gorm:"column:user_id"`
	VideoId  int64 `gorm:"column:video_id"`
	IsDelete bool  `gorm:"column:is_deleted"`
}

func (v Video) TableName() string {
	return "tb_video"
}

type VideoDao struct {
}

var videoDao *VideoDao

func NewVideoDaoInstance() *VideoDao {
	return videoDao
}

func (*VideoDao) AddVideo(video Video) error {
	err := DB.Create(&video).Error
	if err != nil {
		fmt.Println("insert video into db ERROR")
	}
	return err
}

func QueryPublishListByUserId(userId int64) []VideoList {
	var videos []Video
	var userInfo UserInfo

	result := DB.Select("id", "user_id", "play_url", "cover_url", "favorite_count", "comment_count", "title").Where("user_id = ?", userId).Find(&videos)
	DB.First(&userInfo, userId)
	userInfos := SaveUserInfos(userInfo, true)
	n := result.RowsAffected
	videoList := make([]VideoList, n)

	var i int64
	for i = 0; i < n; i++ {
		videoList[i].Id = videos[i].Id
		videoList[i].Author = userInfos
		videoList[i].PlayUrl = videos[i].PlayUrl
		videoList[i].CoverUrl = videos[i].CoverUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].IsFavorite = IsFavorVideo(userId, videos[i].Id)
		videoList[i].Title = videos[i].Title
	}

	return videoList
}
