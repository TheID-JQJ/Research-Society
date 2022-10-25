package model

import (
	"time"

	"gorm.io/gorm"
)

type UserAppoint struct {
	gorm.Model
	UserId          uint            `gorm:"not null" json:"userId"`
	SeatId          uint            `gorm:"not null" json:"seatId"`
	StartTime       time.Time       `gorm:"not null" json:"startTime"`
	EndTIme         time.Time       `gorm:"not null" json:"endTime"`
	Status          string          `gorm:"size:32;not null" json:"status"`
	State           bool            `json:"state"`
	UserInformation UserInformation `gorm:"ForeignKey:UserId"`
	Seat            Seat            `gorm:"ForeignKey:SeatId"`
}
