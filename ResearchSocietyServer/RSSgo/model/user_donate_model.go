package model

import "gorm.io/gorm"

type UserDonate struct {
	gorm.Model
	UserId           uint
	PublicResourceId uint
	State            bool
}
