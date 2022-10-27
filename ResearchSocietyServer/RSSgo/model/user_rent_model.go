package model

import (
	"time"

	"gorm.io/gorm"
)

type UserRent struct {
	gorm.Model
	UserId            uint            `gorm:"not null" json:"userId"`
	PublicResourceId  uint            `gorm:"not null" json:"publicResourceId"`
	RentPrice         float64         `gorm:"not null" json:"rentPrice"`
	RentIntegralPrice int64           `gorm:"not null" json:"rentIntegralPrice"`
	Deposit           float64         `gorm:"not null" json:"deposit"`
	RentTime          time.Time       `gorm:"not null" json:"rentTime"`
	ExpirationTime    time.Time       `gorm:"not null" json:"expirationTime"`
	BackTime          time.Time       `json:"backTime"`
	Status            string          `gorm:"size:32;not null" json:"status"`
	State             bool            `json:"state"`
	UserInformation   UserInformation `gorm:"ForeignKey:UserId"`
	PublicResource    PublicResource  `gorm:"ForeignKey:PublicResourceId"`
}
