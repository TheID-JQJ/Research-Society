package model

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Picture    string
	SeatKindId uint
	State      bool
}
