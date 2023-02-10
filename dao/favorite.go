package dao

import "gorm.io/gorm"

type Favorite struct {
	userId    int64 `gorm:"column:user_id" json:"user_id,omitempty"`
	videoId   int64 `gorm:"column:video_id" json:"video_id,omitempty"`
	isDeleted bool  `gorm:"column:is_deleted" json:"is_deleted"`
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
