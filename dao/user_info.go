package dao

type UserInfo struct {
	ID            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"column:is_follow"`
}

func (u UserInfo) TableName() string {
	return "tb_user_info"
}

func SaveUserInfo(info *UserInfo) {
	DB.Create(info)
}

func QueryUserInfoById(id int64) *UserInfo {
	var userInfo UserInfo
	DB.First(&userInfo, id)
	return &userInfo
}
