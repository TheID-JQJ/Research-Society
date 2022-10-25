package model

import (
	"gorm.io/gorm"
)

type GoodsKinds struct {
	gorm.Model
	Name        string
	Description string
	State       bool
}
