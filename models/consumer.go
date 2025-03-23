package models

import "gorm.io/gorm"

type Consumer struct {
	gorm.Model
	Name string
}
