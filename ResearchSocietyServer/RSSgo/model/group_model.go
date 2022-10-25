package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupNumber     string          `gorm:"size:10;not null;unique" json:"groupNumber"`
	Avatar          string          `gorm:"not null" json:"avatar"`
	GroupName       string          `gorm:"size:32;not null" json:"groupName"`
	OwnerId         uint            `gorm:"not null" json:"ownerId"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:OwnerId"`
}
