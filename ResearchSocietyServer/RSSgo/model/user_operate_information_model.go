package model

import (
	"time"

	"gorm.io/gorm"
)

type UserOperateInformation struct {
	gorm.Model
	Content string
	Time    time.Time
	Result  bool
	State   bool
}
