package dao

import (
	"fmt"
	_ "gorm.io/gorm"
	"time"
)

type Video struct {
	Id            int64     `gorm:"column:id" json:"id,omitempty"`
	UserId        int64     `gorm:"column:user_id" json:"-"`
	Author        UserInfos `gorm:"-" json:"author"`
	PlayUrl       string    `gorm:"column:play_url" json:"play_url,omitempty"`
	CoverUrl      string    `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64     `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64     `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool      `gorm:"-" json:"is_favorite"`
	Title         string    `gorm:"column:title" json:"title"`
	CreatedAt     time.Time `json:"created_at"`
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

func QueryPublishListByUserId(userId int64) []Video {
	var videoList []Video
	var userInfo UserInfo

	result := DB.Select("id", "user_id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "created_at").
		Where("user_id = ?", userId).
		Find(&videoList)
	DB.First(&userInfo, userId)
	userInfos := SaveUserInfos(userInfo, true)
	n := result.RowsAffected

	var i int64

	for i = 0; i < n; i++ {
		videoList[i].Author = userInfos
		videoList[i].IsFavorite = IsFavorVideo(userId, videoList[i].Id)
	}

	return videoList
}

// QueryVideoListByLimitAndTime  返回按投稿时间倒序的视频列表，并限制为最多limit个
func QueryVideoListByLimitAndTime(userId int64, limit int, latestTime time.Time) []Video {
	var videoList []Video

	result := DB.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at ASC").Limit(limit).
		Select([]string{"id", "user_id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "created_at"}).
		Find(videoList)
	n := result.RowsAffected

	var i int64
	for i = 0; i < n; i++ {
		var userInfo UserInfo
		DB.Where("id = ?", videoList[i].UserId).Find(&userInfo)
		flag := IsFollow(userId, videoList[i].UserId)
		userInfos := SaveUserInfos(userInfo, flag)

		videoList[i].Author = userInfos
		videoList[i].IsFavorite = IsFavorVideo(userId, videoList[i].Id)
	}

	return videoList
}
