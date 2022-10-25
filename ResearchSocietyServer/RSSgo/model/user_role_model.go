package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId uint
	Role   string
	State  bool
}
