package model

import "gorm.io/gorm"

type UserPassword struct {
	gorm.Model
	UserId      uint
	UserNumber  string
	Email       string
	PhoneNumber string
	Password    string
	State       bool
}
