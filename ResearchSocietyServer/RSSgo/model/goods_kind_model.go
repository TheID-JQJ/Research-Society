package model

import (
	"gorm.io/gorm"
)

type GoodsKind struct {
	gorm.Model
	Name        string `gorm:"size:32;not null;unique" json:"name"`
	Description string `gorm:"not null" json:"description"`
	State       bool   `json:"state"`
}
