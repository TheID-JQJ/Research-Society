package model

import "gorm.io/gorm"

type UserRealName struct {
	gorm.Model
	UserId          uint            `gorm:"not null;unique" json:"userId"`
	Name            string          `gorm:"size:32;not null" json:"name"`
	IdNumber        string          `gorm:"size:18;not null;unique" json:"idNumber"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
