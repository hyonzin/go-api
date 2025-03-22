package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type record struct {
	ID       string `json:"id"`
	Datetime string `json:"datetime"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

var records = []record{
	{ID: "1", Datetime: "2025-03-22 18:37:00", Title: "Learning go - 1 day", Content: "rest api with gin"},
	{ID: "2", Datetime: "2025-03-22 18:38:00", Title: "dummy", Content: "just a record"},
	{ID: "3", Datetime: "2025-03-22 18:39:00", Title: "whatever", Content: "ya"},
}

func main() {
	router := gin.Default()
	router.GET("/records", getRecords)

	router.Run("0.0.0.0:30002")
}

func getRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, records)
}
