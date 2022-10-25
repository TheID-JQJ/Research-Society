package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAppoint struct {
	gorm.Model
	UserId    uint
	SeatId    uint
	StartTime time.Time
	EndTIme   time.Time
	Status    string
	State     bool
}
