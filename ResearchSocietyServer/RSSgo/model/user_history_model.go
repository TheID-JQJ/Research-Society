package model

import (
	"time"

	"gorm.io/gorm"
)

type UserHistory struct {
	gorm.Model
	UserId           uint             `gorm:"not null" json:"userId"`
	GoodsId          uint             `gorm:"not null" json:"goodsId"`
	HistoryTime      time.Time        `gorm:"not null" json:"historyTime"`
	State            bool             `json:"state"`
	UserInformation  UserInformation  `gorm:"ForeignKey:UserId"`
	GoodsInformation GoodsInformation `gorm:"ForeignKey:GoodsId"`
}
