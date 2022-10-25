package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupNumber string
	Avatar      string
	GroupName   string
	OwnerId     uint
	State       bool
}
