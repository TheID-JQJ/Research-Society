package model

import "gorm.io/gorm"

type FinanceKind struct {
	gorm.Model
	Name  string
	State bool
}
