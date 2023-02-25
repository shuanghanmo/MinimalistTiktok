package dao

type UserInfo struct {
	ID            int64  `gorm:"column:id" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"-" json:"is_follow"`
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
