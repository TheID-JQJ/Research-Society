package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"type:varchar(16);not null" json:"password"`
}
