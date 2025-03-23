package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/controllers"
)

func MessageQueueRoutes(r *gin.Engine) {
	userGroup := r.Group("/mq")
	{
		userGroup.POST("/produce", controllers.ProduceMessage)
		userGroup.POST("/consume", controllers.ConsumeMessage)
	}
}
