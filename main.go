package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type record struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdat"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

var records = []record{
	{ID: "1", CreatedAt: "2025-03-22 18:37:00", Title: "Learning go - 1 day", Content: "rest api with gin"},
	{ID: "2", CreatedAt: "2025-03-22 18:38:00", Title: "dummy", Content: "just a record"},
	{ID: "3", CreatedAt: "2025-03-22 18:39:00", Title: "whatever", Content: "ya"},
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)
	router.POST("/records", postRecords)

	router.Run("0.0.0.0:30002")
}

func getRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, records)
}

func postRecords(c *gin.Context) {
	var newRecord record

	// Validate input record
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}

	// Add record
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}
