package service

import (
	"hkc.ink/rss/database"
	"hkc.ink/rss/dto"

	"gorm.io/gorm"
)

type UserService struct {
}

var userDao database.UserDao

// 用户注册
func (UserService) Register(db *gorm.DB, rm *dto.RegisterModel) error {
	err := userDao.Insert(db, rm.ToUserModel())
	return err
}
