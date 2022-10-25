package model

import "gorm.io/gorm"

type SeatKind struct {
	gorm.Model
	Name  string
	State bool
}
