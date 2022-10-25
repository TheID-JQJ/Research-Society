package model

import (
	"time"

	"gorm.io/gorm"
)

type UserOrders struct {
	gorm.Model
	UserId            uint
	GoodsId           uint
	OrderNumber       string
	GoodsQuantity     int64
	RealPrice         float64
	RealIntegralPrice int64
	Validity          time.Time
	CreateTime        time.Time
	PayTime           time.Time
	ConsumeTime       time.Time
	Status            string
	State             bool
}
