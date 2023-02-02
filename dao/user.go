package dao

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64  `gorm:"column:id"`
	UserName  string `gorm:"column:user_name"`
	PassWord  string `gorm:"column:pass_word"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (u User) TableName() string {
	return "tb_user"
}

func SaveUser(user *User) {
	DB.Create(user)
}

func QueryByUser(userName string, passWord string) *User {
	var user User
	DB.Where("user_name = ? and pass_word = ?", userName, passWord).Find(&user).Limit(1)
	return &user
}

func QueryByUserName(userName string) *User {
	var user User
	DB.Where("user_name = ?", userName).Find(&user).Limit(1)
	return &user
}

func UpdateUser(user *User) {
	DB.Model(&User{}).Updates(user)
}

func QueryById(id int64) *User {
	var user User
	DB.First(&user, id)
	return &user
}

func DeleteUser(id int64) {
	DB.Where("id = ?", id).Delete(&User{})
}
