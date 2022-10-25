package model

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	gorm.Model
	KindId          uint            `gorm:"not null" json:"kindId"`
	UserId          uint            `gorm:"not null" json:"userId"`
	Time            time.Time       `gorm:"not null" json:"time"`
	Money           int64           `gorm:"not null" json:"money"`
	Details         string          `gorm:"not null" json:"details"`
	State           bool            `json:"state"`
	FinanceKind     FinanceKind     `gorm:"ForeignKey:KindId"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
