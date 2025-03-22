package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/initializers"
	"github.com/hyonzin/go-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	// Record Routes
	routes.RecordRoutes(r)

	r.Run("0.0.0.0:30002")
}
