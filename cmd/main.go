package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-message-queue-broker/db"
	"github.com/hyonzin/go-message-queue-broker/initializers"
	"github.com/hyonzin/go-message-queue-broker/routes"
)

func init() {
	initializers.LoadEnvVariables()
	db.ConnectDB()
}

func main() {
	r := gin.Default()

	// Record Routes
	routes.RecordRoutes(r)

	// MQ Routes
	routes.MessageQueueRoutes(r)

	r.Run("0.0.0.0:30002")
}
