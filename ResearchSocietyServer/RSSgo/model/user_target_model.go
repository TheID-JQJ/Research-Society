package model

import "gorm.io/gorm"

type UserTarget struct {
	gorm.Model
	UserId          uint            `gorm:"not null;unique" json:"userId"`
	TypeId          uint            `gorm:"not null" json:"typeId"`
	Details         string          `json:"details"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
	TargetType      TargetType      `gorm:"ForeignKey:TypeId"`
}
