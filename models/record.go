package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Title   string
	Content string
}
