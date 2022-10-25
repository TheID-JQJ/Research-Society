package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSeat struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	SeatId          uint            `gorm:"not null" json:"seatId"`
	StartTime       time.Time       `gorm:"not null" json:"startTime"`
	EndTime         time.Time       `json:"endTime"`
	Status          string          `gorm:"size:32;not null" json:"status"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
	Seat            Seat            `gorm:"ForeignKey:SeatId"`
}
