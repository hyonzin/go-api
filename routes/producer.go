package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/controllers"
)

func ProducerRoutes(r *gin.Engine) {
	userGroup := r.Group("/producers")
	{
		// Producer Routes
		userGroup.POST("", controllers.ProducersCreate)       // Create
		userGroup.GET("", controllers.ProducersIndex)         // Read all
		userGroup.GET("/:id", controllers.ProducersShow)      // Read One
		userGroup.PUT("/:id", controllers.ProducersUpdate)    // Update
		userGroup.DELETE("/:id", controllers.ProducersDelete) // Delete
	}
}
