package dao

import (
	"fmt"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	UserId        int64  `gorm:"column:user_id"`
	PlayUrl       string `gorm:"column:play_url" json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count,omitempty"`
	//IsFavorite    bool   `gorm:"column:is_favorite" json:"is_favorite,omitempty"`
	Title string `gorm:"column:title"`
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
		videoList[i].IsFavorite = IsFavorVideo(userId, videos[i].Id)
		videoList[i].Title = videos[i].Title
	}

	return videoList
}

// PlusOneFavorByUserIdAndVideoId 增加一个赞
func PlusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE tb_video SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		//功能未完善，需要先查询数据库中是否有记录，有就将is_deleted设为1，没有则插入
		if err := tx.Exec("INSERT INTO `user_favor_videos` (`user_id`,`video_id`,'is_deleted') VALUES (?,?,0)", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// MinusOneFavorByUserIdAndVideoId 减少一个赞
func MinusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		//执行-1之前需要先判断是否合法（不能被减少为负数
		if err := tx.Exec("UPDATE tb_video SET favorite_count = favorite_count - 1 WHERE id = ? AND favorite_count > 0", videoId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_favor_videos SET is_deleted = 1 WHERE `user_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

func IsFavorVideo(userId int64, videoId int64) bool {
	var count int64
	DB.Table("user_favor_videos").Where("`user_id` = ? AND `video_id` = ? AND `is_deleted` = 0", userId, videoId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
