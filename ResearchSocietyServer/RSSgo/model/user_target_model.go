package model

import "gorm.io/gorm"

type UserTarget struct {
	gorm.Model
	UserId  uint
	TypeId  uint
	Details string
	State   bool
}
