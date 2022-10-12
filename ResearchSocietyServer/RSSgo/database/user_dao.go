package database

import (
	"gorm.io/gorm"
	"hkc.ink/rss/model"
)

type UserDao struct {
}

// 用户插入
func (UserDao) Insert(db *gorm.DB, user *model.User) error {
	result := db.Create(user)
	return result.Error
}
