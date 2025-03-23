package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/initializers"
	"github.com/hyonzin/go-api/models"
)

func ConsumersCreate(c *gin.Context) {
	// Get data from req body
	var body struct {
		Name string
	}
	c.Bind(&body)

	// Create a consumer
	consumer := models.Consumer{Name: body.Name}
	result := initializers.DB.Create(&consumer)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"consumer": consumer,
	})
}

func ConsumersIndex(c *gin.Context) {
	// Get all the consumers
	var consumers []models.Consumer
	initializers.DB.Find(&consumers)

	// Return consumers in response
	c.JSON(200, gin.H{
		"consumers": consumers,
	})
}

func ConsumersShow(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Get a sing consumer
	var consumer models.Consumer
	initializers.DB.First(&consumer, id)

	// Return consumer in response
	c.JSON(200, gin.H{
		"consumer": consumer,
	})
}

func ConsumersUpdate(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Name string
	}
	c.Bind(&body)

	// Get a single consumer that we want to update
	var consumer models.Consumer
	initializers.DB.First(&consumer, id)

	// Update it
	initializers.DB.Model(&consumer).Updates(models.Consumer{
		Name: body.Name,
	})

	// Return response
	c.JSON(200, gin.H{
		"consumer": consumer,
	})
}

func ConsumersDelete(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Delete the Consumer
	initializers.DB.Delete(&models.Consumer{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Consumer removed Successfully",
	})
}
