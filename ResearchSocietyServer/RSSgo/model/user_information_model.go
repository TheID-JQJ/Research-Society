package model

import "gorm.io/gorm"

type UserInformation struct {
	gorm.Model
	UserNumber   string
	Avatar       string
	Username     string
	Sex          string
	Introduction string
	Email        string
	PhoneNumber  string
	State        bool
}
