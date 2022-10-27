package model

import (
	"gorm.io/gorm"
)

type GoodsKind struct {
	gorm.Model
	Name  string `gorm:"size:32;not null;unique" json:"name"`
	State bool   `json:"state"`
}
