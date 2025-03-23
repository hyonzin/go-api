package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/hyonzin/go-message-queue-broker/db"
	"github.com/hyonzin/go-message-queue-broker/models"
)

func ProduceMessage(c *gin.Context) {
	// Get data from req body
	var body struct {
		Topic    string
		Contents string
	}
	c.Bind(&body)

	// Create a message
	message := models.Message{Topic: body.Topic, Contents: body.Contents}
	result := db.DB.Create(&message)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"message": message,
	})
}

func ConsumeMessage(c *gin.Context) {
	// Get data from req body
	var body struct {
		Topic string
	}
	c.Bind(&body)

	// Get a sing consumer
	var message models.Message
	result := db.DB.First(&message, body)

	if result.Error != nil {
		c.JSON(200, gin.H{
			"topic":       body.Topic,
			"is_consumed": false,
			"message":     nil,
		})
		return
	}

	// Delete the Consumer
	db.DB.Delete(&models.Message{}, message)

	// Return consumer in response
	c.JSON(200, gin.H{
		"topic":       body.Topic,
		"is_consumed": true,
		"message":     message,
	})
}
