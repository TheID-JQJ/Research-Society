package model

import "gorm.io/gorm"

type GoodsInformation struct {
	gorm.Model
	Picture       string    `gorm:"not null" json:"picture"`
	Name          string    `gorm:"size:32;not null" json:"name"`
	Price         float64   `gorm:"not null" json:"price"`
	IntegralPrice int64     `gorm:"not null" json:"integralPrice"`
	MonthlySales  int64     `gorm:"default:0;not null" json:"monthlySales"`
	Inventory     int64     `gorm:"default:0;not null" json:"inventory"`
	KindId        uint      `gorm:"not null" json:"kindId"`
	State         bool      `json:"state"`
	GoodsKind     GoodsKind `gorm:"ForeignKey:KindId"`
}
