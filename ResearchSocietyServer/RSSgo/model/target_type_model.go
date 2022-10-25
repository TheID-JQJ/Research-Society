package model

import "gorm.io/gorm"

type TargetType struct {
	gorm.Model
	Name  string
	State bool
}
