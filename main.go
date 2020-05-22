package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var events = make(map[string]string)

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func getEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

func createEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

func updateEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

func main() {
	r := gin.Default()

	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/events", createEvent)
	r.PUT("/events/:id", updateEvent)

	r.Run(":8080")
}
