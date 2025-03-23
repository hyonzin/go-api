package main

import (
	"github.com/hyonzin/go-message-queue-broker/db"
	"github.com/hyonzin/go-message-queue-broker/initializers"
	"github.com/hyonzin/go-message-queue-broker/models"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Record{})
	db.DB.AutoMigrate(&models.Producer{})
	db.DB.AutoMigrate(&models.Consumer{})
	db.DB.AutoMigrate(&models.Message{})
}
