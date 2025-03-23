package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-message-queue-broker/controllers"
)

func RecordRoutes(r *gin.Engine) {
	userGroup := r.Group("/records")
	{
		// Record Routes
		userGroup.POST("", controllers.RecordsCreate)       // Create
		userGroup.GET("", controllers.RecordsIndex)         // Read all
		userGroup.GET("/:id", controllers.RecordsShow)      // Read One
		userGroup.PUT("/:id", controllers.RecordsUpdate)    // Update
		userGroup.DELETE("/:id", controllers.RecordsDelete) // Delete
	}
}
