package dao

import "gorm.io/gorm"

type Relation struct {
	UserId    int64 `gorm:"column:user_id" json:"user_id,omitempty"`
	ToUserId  int64 `gorm:"column:to_user_id" json:"to_user_id,omitempty"`
	IsDeleted bool  `gorm:"column:is_deleted" json:"is_deleted"`
}

func (r Relation) TableName() string {
	return "tb_relation"
}

func FollowAction(userId int64, toUserId int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE tb_user_info SET follow_count = follow_count + 1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE tb_user_info SET follower_count = follower_count + 1 WHERE id = ?", toUserId).Error; err != nil {
			return err
		}
		var count int64
		tx.Table("tb_relation").Where("user_id = ? and to_user_id = ?", userId, toUserId).Count(&count)
		if count == 0 {
			if err := tx.Exec("INSERT INTO `tb_relation` (`user_id`,`to_user_id`,`is_deleted`) VALUES (?,?,0)", userId, toUserId).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Exec("UPDATE tb_relation SET is_deleted = 0 WHERE `user_id` = ? AND `to_user_id` = ?", userId, toUserId).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func UnFollowAction(userId int64, toUserId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		//执行-1之前需要先判断是否合法（不能被减少为负数
		if err := tx.Exec("UPDATE tb_user_info SET follow_count = follow_count - 1 WHERE id = ? AND follow_count > 0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE tb_user_info SET follower_count = follower_count - 1 WHERE id = ? AND follower_count > 0", toUserId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE tb_relation SET is_deleted = 1 WHERE `user_id` = ? AND `to_user_id` = ?", userId, toUserId).Error; err != nil {
			return err
		}
		return nil
	})
}

func IsFollow(userId int64, toUserId int64) bool {
	var count int64
	DB.Table("tb_relation").Where("`user_id` = ? AND `to_user_id` = ? AND `is_deleted` = 0", userId, toUserId).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
