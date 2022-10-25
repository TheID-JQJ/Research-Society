package model

import (
	"time"

	"gorm.io/gorm"
)

type UserOrders struct {
	gorm.Model
	UserId            uint             `gorm:"not null" json:"userId"`
	GoodsId           uint             `gorm:"not null" json:"goodsId"`
	OrderNumber       string           `gorm:"size:18;not null;unique" json:"orderNumber"`
	Quantity          int64            `gorm:"not null" json:"quantity"`
	RealPrice         float64          `gorm:"not null" json:"realPrice"`
	RealIntegralPrice int64            `gorm:"not null" json:"realIntegralPrice"`
	Validity          time.Time        `json:"validity"`
	CreateTime        time.Time        `gorm:"not null" json:"createTime"`
	PayTime           time.Time        `json:"payTime"`
	ConsumeTime       time.Time        `json:"consumeTime"`
	Status            string           `gorm:"size:32;not null" json:"status"`
	State             bool             `json:"state"`
	UserInformation   UserInformation  `gorm:"ForeignKey:UserId"`
	GoodsInformation  GoodsInformation `gorm:"ForeignKey:GoodsId"`
}
