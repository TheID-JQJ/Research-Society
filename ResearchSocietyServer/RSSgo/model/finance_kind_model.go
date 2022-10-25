package model

import "gorm.io/gorm"

type FinanceKind struct {
	gorm.Model
	Name  string `gorm:"not null;unique" json:"name"`
	State bool   `json:"state"`
}
