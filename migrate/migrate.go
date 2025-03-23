package main

import (
	"github.com/hyonzin/go-api/initializers"
	"github.com/hyonzin/go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Record{})
	initializers.DB.AutoMigrate(&models.Producer{})
	initializers.DB.AutoMigrate(&models.Consumer{})
	initializers.DB.AutoMigrate(&models.Message{})
}
