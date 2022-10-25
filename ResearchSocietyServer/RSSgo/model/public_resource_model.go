package model

import "gorm.io/gorm"

type PublicResource struct {
	gorm.Model
	Picture            string             `gorm:"not null" json:"picture"`
	Name               string             `gorm:"size:32;not null" json:"name"`
	Price              float64            `gorm:"not null" json:"price"`
	IntegralPrice      int64              `gorm:"not null" json:"integralPrice"`
	RentAmount         int64              `gorm:"not null;default:0" json:"rentAmount"`
	KindId             uint               `gorm:"not null" json:"kindId"`
	State              bool               `json:"state"`
	PublicResourceKind PublicResourceKind `gorm:"ForeignKey:KindId"`
}
