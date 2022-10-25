package model

import "gorm.io/gorm"

type UserAccount struct {
	gorm.Model
	UserId   uint
	Residual float64
	Integral int64
	Duration int64
	State    bool
}
