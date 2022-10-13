package database

import (
	"gorm.io/gorm"
	"hkc.ink/rss/model"
)

type UserDao struct {
}

// 查询
func (UserDao) GetByID(db *gorm.DB, id uint) (model.User, error) {
	user := model.User{}
	result := db.First(&user, id)
	return user, result.Error
}
func (UserDao) GetAll(db *gorm.DB) ([]model.User, error) {
	users := []model.User{}
	result := db.Find(&users)
	return users, result.Error
}

// 根据邮箱查询用户
func (UserDao) GetByEmail(db *gorm.DB, email string) (model.User, error) {
	user := model.User{}
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

// 插入
func (UserDao) Insert(db *gorm.DB, user *model.User) (uint, error) {
	result := db.Create(user)
	return user.ID, result.Error
}

// 删除
func (ud UserDao) DeleteByID(db *gorm.DB, id uint) error {
	if _, err := ud.GetByID(db, id); err != nil {
		return err
	}
	db.Delete(&model.User{}, id)
	return nil
}

// 修改
func (ud UserDao) Update(db *gorm.DB, user *model.User) error {
	if _, err := ud.GetByID(db, user.ID); err != nil {
		return err
	}
	db.Model(&user).Where("id = ?", user.ID).Updates(*user)
	return nil
}
