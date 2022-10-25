package model

import "gorm.io/gorm"

type UserStar struct {
	gorm.Model
	UserId  uint
	GoodsId uint
	State   bool
}
