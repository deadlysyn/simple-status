package main

import (
	"github.com/gin-gonic/gin"
)

// Event describes a status event
type Event struct {
	Status      string `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
	Updated     string `json:"updated,omitempty"`
}

// Events maps service names to Events
var Events = make(map[string]Event)

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
