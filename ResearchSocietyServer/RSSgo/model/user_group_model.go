package model

import "gorm.io/gorm"

type UserGroup struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	GroupId         uint            `gorm:"not null" json:"groupId"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
	Group           Group           `gorm:"ForeignKey:GroupId"`
}
