package model

import "gorm.io/gorm"

type PublicResourceKind struct {
	gorm.Model
	Name  string
	State bool
}
