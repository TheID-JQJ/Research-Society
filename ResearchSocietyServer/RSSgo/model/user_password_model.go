package model

import "gorm.io/gorm"

type UserPassword struct {
	gorm.Model
	UserId          uint            `gorm:"not null;unique" json:"userId"`
	Password        string          `gorm:"not null" json:"password"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
