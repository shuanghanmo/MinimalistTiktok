package dao

import "gorm.io/gorm"

type Favorite struct {
	UserId    int64 `gorm:"column:user_id" json:"user_id,omitempty"`
	VideoId   int64 `gorm:"column:video_id" json:"video_id,omitempty"`
	IsDeleted bool  `gorm:"column:is_deleted" json:"is_deleted"`
}

func (f Favorite) TableName() string {
	return "user_favor_videos"
}

// PlusOneFavorByUserIdAndVideoId 增加一个赞
func PlusOneFavorByUserIdAndVideoId(userId int64, videoId int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE tb_video SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
			return err
		}
		var count int64
		tx.Table("user_favor_videos").Where("user_id = ? and video_id = ?", userId, videoId).Count(&count)
		if count == 0 {
			if err := tx.Exec("INSERT INTO `user_favor_videos` (`user_id`,`video_id`,`is_deleted`) VALUES (?,?,0)", userId, videoId).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Exec("UPDATE user_favor_videos SET is_deleted = 0 WHERE `user_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
				return err
			}
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

// 通过userId查询点赞视频列表
func QueryFavorVideosByUserId(userId int64) []VideoList {
	var favorites []Favorite
	var video Video
	var userInfo UserInfo

	result := DB.Select("video_id").Where("`user_id` = ? and `is_deleted` = 0", userId).Find(&favorites)

	n := result.RowsAffected
	videoList := make([]VideoList, n)

	var i int64
	for i = 0; i < n; i++ {
		DB.Select("id", "user_id", "play_url", "cover_url", "favorite_count", "comment_count", "title").Where("id = ?", favorites[i].VideoId).Find(&video)
		DB.First(&userInfo, video.UserId)

		videoList[i].Id = video.Id
		videoList[i].Author = userInfo
		videoList[i].PlayUrl = video.PlayUrl
		videoList[i].CoverUrl = video.CoverUrl
		videoList[i].FavoriteCount = video.FavoriteCount
		videoList[i].CommentCount = video.CommentCount
		videoList[i].IsFavorite = true
		videoList[i].Title = video.Title
	}

	return videoList
}
