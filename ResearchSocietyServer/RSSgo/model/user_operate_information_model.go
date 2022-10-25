package model

import (
	"time"

	"gorm.io/gorm"
)

type UserOperateInformation struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	Content         string          `gorm:"not null" json:"content"`
	Time            time.Time       `gorm:"not null" json:"time"`
	Result          bool            `gorm:"not null" json:"result"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
