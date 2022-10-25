package model

import "gorm.io/gorm"

type UserInformation struct {
	gorm.Model
	UserNumber   string `gorm:"size:10;not null;unique" json:"userNumber"`
	Avatar       string `gorm:"not null" json:"avatar"`
	Username     string `gorm:"size:32;not null" json:"username"`
	Sex          string `gorm:"size:2;not null" json:"sex"`
	Introduction string `gorm:"not null" json:"introduction"`
	Email        string `gorm:"size:255;not null;unique" json:"email"`
	PhoneNumber  string `gorm:"size:11;not null;unique" json:"phoneNumber"`
	State        bool   `json:"state"`
}
