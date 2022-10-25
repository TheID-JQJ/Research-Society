package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	UserId  uint
	Content string
	State   bool
}
