package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	Content         string          `gorm:"not null" json:"content"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
}
