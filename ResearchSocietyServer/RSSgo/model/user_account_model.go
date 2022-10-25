package model

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	UserId          uint            `gorm:"not null;unique" json:"userId"`
	Residual        float64         `gorm:"not null" json:"residual"`
	Integral        int64           `gorm:"not null" json:"integral"`
	Duration        int64           `gorm:"not null" json:"duration"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
