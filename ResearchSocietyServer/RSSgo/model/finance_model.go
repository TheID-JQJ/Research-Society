package model

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	gorm.Model
	Kind    string
	UserId  uint
	Time    time.Time
	Money   int64
	Details string
	State   bool
}
