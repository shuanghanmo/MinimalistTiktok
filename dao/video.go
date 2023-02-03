package dao

import (
	_ "gorm.io/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	UserId        int64  `gorm:"column:user_id"`
	PlayUrl       string `gorm:"column:play_url" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"column:is_favorite" json:"is_favorite,omitempty"`
	Title         string `gorm:"column:title"`
}

type VideoList struct {
	Id            int64    `gorm:"column:id" json:"id,omitempty"`
	Author        UserInfo `gorm:"column:author" json:"author"`
	PlayUrl       string   `gorm:"column:play_url" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64    `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64    `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool     `gorm:"column:is_favorite" json:"is_favorite,omitempty"`
	Title         string   `gorm:"column:title" json:"title"`
}

func (v Video) TableName() string {
	return "tb_video"
}

func QueryPublishListByUserId(userId int64) []VideoList {
	var videos []Video
	var userInfo UserInfo

	result := DB.Where("user_id = ?", userId).Find(&videos)
	DB.First(&userInfo, userId)
	n := result.RowsAffected
	videoList := make([]VideoList, n)

	var i int64
	for i = 0; i < n; i++ {
		videoList[i].Id = videos[i].Id
		videoList[i].Author = userInfo
		videoList[i].PlayUrl = videos[i].PlayUrl
		videoList[i].CoverUrl = videos[i].CoverUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].IsFavorite = videos[i].IsFavorite
		videoList[i].Title = videos[i].Title
	}

	return videoList
}
