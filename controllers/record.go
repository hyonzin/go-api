package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyonzin/go-api/initializers"
	"github.com/hyonzin/go-api/models"
)

func RecordsCreate(c *gin.Context) {
	// Get data from req body
	var body struct {
		Title   string
		Content string
	}
	c.Bind(&body)

	// Create a record
	record := models.Record{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&record)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"record": record,
	})
}

func RecordsIndex(c *gin.Context) {
	// Get all the records
	var records []models.Record
	initializers.DB.Find(&records)

	// Return records in response
	c.JSON(200, gin.H{
		"records": records,
	})
}

func RecordsShow(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Get a sing record
	var record models.Record
	initializers.DB.First(&record, id)

	// Return record in response
	c.JSON(200, gin.H{
		"record": record,
	})
}

func RecordsUpdate(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// get the data of req body
	var body struct {
		Title   string
		Content string
	}
	c.Bind(&body)

	// Get a single record that we want to update
	var record models.Record
	initializers.DB.First(&record, id)

	// Update it
	initializers.DB.Model(&record).Updates(models.Record{
		Title:   body.Title,
		Content: body.Content,
	})

	// Return response
	c.JSON(200, gin.H{
		"record": record,
	})
}

func RecordsDelete(c *gin.Context) {
	// Get id from URL param
	id := c.Param("id")

	// Delete the Record
	initializers.DB.Delete(&models.Record{}, id)

	// Return response
	c.JSON(200, gin.H{
		"message": "Record removed Successfully",
	})
}
