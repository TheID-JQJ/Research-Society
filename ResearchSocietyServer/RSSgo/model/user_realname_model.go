package model

import "gorm.io/gorm"

type UserRealName struct {
	gorm.Model
	UserId   uint
	Name     string
	IdNumber string
	State    bool
}
