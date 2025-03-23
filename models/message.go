package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Topic    string
	Contents string
}
