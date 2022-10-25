package model

import "gorm.io/gorm"

type UserShoppingCart struct {
	gorm.Model
	UserId           uint             `gorm:"not null" json:"userId"`
	GoodsId          uint             `gorm:"not null" json:"goodsId"`
	State            bool             `json:"state"`
	UserInformation  UserInformation  `gorm:"ForeignKey:UserId"`
	GoodsInformation GoodsInformation `gorm:"ForeignKey:GoodsId"`
}
