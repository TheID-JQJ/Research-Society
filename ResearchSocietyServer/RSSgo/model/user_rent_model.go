package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRent struct {
	gorm.Model
	UserId            uint
	PublicResourceId  uint
	RentPrice         float64
	RentIntegralPrice int64
	Deposit           float64
	RentTime          time.Time
	BackTime          time.Time
	Status            string
	State             bool
}
