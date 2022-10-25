package model

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Picture  string   `gorm:"not null" json:"picture"`
	KindId   uint     `gorm:"not null" json:"seatKindId"`
	State    bool     `json:"state"`
	SeatKind SeatKind `gorm:"ForeignKey:KindId"`
}
