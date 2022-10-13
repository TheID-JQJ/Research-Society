package service

import (
	"log"

	"hkc.ink/rss/database"
	"hkc.ink/rss/dto"
	"hkc.ink/rss/model"

	"gorm.io/gorm"
)

type UserService struct {
	database.UserDao
}

// 用户注册
func (userService UserService) Register(db *gorm.DB, rd *dto.RegisterDto) (uint, error) {
	id, err := userService.Insert(db, rd.ToUserModel())
	return id, err
}

// 邮箱是否存在
func (userService UserService) ExistEmail(db *gorm.DB, email string) bool {
	//查询email
	result := db.Unscoped().Where("email = ?", email).First(&model.User{})
	if result.Error != nil {
		log.Println("sql执行错误：" + result.Error.Error())
		return false
	}
	return true
}
