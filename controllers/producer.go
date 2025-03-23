package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/initializers"
	"github.com/hyonzin/go-api/models"
)

func ProducersCreate(c *gin.Context) {
	// Get data from req body
	var body struct {
		Name string
	}
	c.Bind(&body)

	// Create a producer
	producer := models.Producer{Name: body.Name}
	result := initializers.DB.Create(&producer)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"producer": producer,
	})
}

func ProducersIndex(c *gin.Context) {
	// Get all the producers
	var producers []models.Producer
	initializers.DB.Find(&producers)

	// Return producers in response
	c.JSON(200, gin.H{
		"producers": producers,
	})
}

func ProducersShow(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Get a sing producer
	var producer models.Producer
	initializers.DB.First(&producer, id)

	// Return producer in response
	c.JSON(200, gin.H{
		"producer": producer,
	})
}

func ProducersUpdate(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Name string
	}
	c.Bind(&body)

	// Get a single producer that we want to update
	var producer models.Producer
	initializers.DB.First(&producer, id)

	// Update it
	initializers.DB.Model(&producer).Updates(models.Producer{
		Name: body.Name,
	})

	// Return response
	c.JSON(200, gin.H{
		"producer": producer,
	})
}

func ProducersDelete(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Delete the Producer
	initializers.DB.Delete(&models.Producer{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Producer removed Successfully",
	})
}
