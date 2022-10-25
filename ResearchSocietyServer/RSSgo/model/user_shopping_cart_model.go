package model

import "gorm.io/gorm"

type UserShoppingCart struct {
	gorm.Model
	UserId  uint
	GoodsId uint
	State   bool
}
