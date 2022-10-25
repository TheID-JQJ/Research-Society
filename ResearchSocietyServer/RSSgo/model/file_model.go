package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName string
	FilePath string
}
