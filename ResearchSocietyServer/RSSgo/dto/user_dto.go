package dto

import (
	"gorm.io/gorm"
	"hkc.ink/rss/model"
)

type UserDto struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" binding:"required,email" required_msg:"邮箱不能为空" email_msg:"邮箱格式不正确"`
	Password string `json:"password" binding:"required" required_msg:"密码不能为空"`
}

func (ud UserDto) ToUserModel() *model.User {
	return &model.User{Model: gorm.Model{ID: ud.ID}, Email: ud.Email, Password: ud.Password}
}

type RegisterDto struct {
	Email    string `json:"email" binding:"required,email" required_msg:"邮箱不能为空" email_msg:"邮箱格式不正确"`
	Password string `json:"password" binding:"required" required_msg:"密码不能为空"`
}

func (rd RegisterDto) ToUserModel() *model.User {
	return &model.User{Email: rd.Email, Password: rd.Password}
}
