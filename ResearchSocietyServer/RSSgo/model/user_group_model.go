package model

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	UserId  uint
	GroupId uint
	State   bool
}
