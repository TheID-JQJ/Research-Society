package model

import (
	"time"

	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Title   string
	Content string
	Time    time.Time
	State   bool
}
