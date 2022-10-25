package model

import "gorm.io/gorm"

type UserDonate struct {
	gorm.Model
	UserId           uint            `gorm:"not null" json:"userId"`
	PublicResourceId uint            `gorm:"not null;unique" json:"publicResourceId"`
	State            bool            `json:"state"`
	UserInformation  UserInformation `gorm:"ForeignKey:UserId"`
	PublicResource   PublicResource  `gorm:"ForeignKey:PublicResourceId"`
}
