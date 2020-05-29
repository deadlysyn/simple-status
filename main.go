package main

import (
	"github.com/gin-gonic/gin"
)

// Event fields require when POSTing or creating new events
type Event struct {
	Service     string `json:"service" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	Updated     string `json:"updated"`
}

// EventData fields required when PUTting or updating events
type EventData struct {
	Status      string `json:"status,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:"required"`
	Updated     string `json:"updated,omitempty"`
}

// Events maps service names to data about status events
var Events = make(map[string]EventData)

func main() {
	router := gin.Default()

	router.Static("/public", "./public")
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/html/index.html")
	})

	api := router.Group("/api")
	api.GET("/events", GetEvents)
	api.GET("/events/:service", GetEvent)
	api.POST("/events", CreateEvent)
	api.PUT("/events/:service", UpdateEvent)

	router.Run()
}
