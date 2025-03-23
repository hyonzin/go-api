package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/controllers"
)

func ConsumerRoutes(r *gin.Engine) {
	userGroup := r.Group("/consumers")
	{
		// Consumer Routes
		userGroup.POST("", controllers.ConsumersCreate)       // Create
		userGroup.GET("", controllers.ConsumersIndex)         // Read all
		userGroup.GET("/:id", controllers.ConsumersShow)      // Read One
		userGroup.PUT("/:id", controllers.ConsumersUpdate)    // Update
		userGroup.DELETE("/:id", controllers.ConsumersDelete) // Delete
	}
}
