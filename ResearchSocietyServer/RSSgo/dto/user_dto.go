package dto

import (
	"gorm.io/gorm"
	"hkc.ink/rss/model"
)

type UserDto struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ud UserDto) ToUserModel() *model.User {
	return &model.User{Model: gorm.Model{ID: ud.ID}, Email: ud.Email, Password: ud.Password}
}

type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rd RegisterDto) ToUserModel() *model.User {
	return &model.User{Email: rd.Email, Password: rd.Password}
}
