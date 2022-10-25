package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	Role            string          `gorm:"size:32;not null" json:"role"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
