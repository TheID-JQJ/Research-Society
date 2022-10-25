package model

import (
	"time"

	"gorm.io/gorm"
)

type Notice struct {
	gorm.Model
	Title   string    `gorm:"size:255;not null" json:"title"`
	Content string    `gorm:"not null" json:"content"`
	Time    time.Time `gorm:"not null" json:"time"`
	State   bool      `json:"state"`
}
