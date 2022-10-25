package model

import (
	"time"

	"gorm.io/gorm"
)

type UserHistory struct {
	gorm.Model
	UserId      uint
	GoodsId     uint
	HistoryTime time.Time
	State       bool
}
