package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSeat struct {
	gorm.Model
	UserId    uint
	SeatId    uint
	StartTime time.Time
	EndTime   time.Time
	State     bool
}
