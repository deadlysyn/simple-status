package main

import (
	"github.com/gin-gonic/gin"
)

// Event describes a "status event"
type Event struct {
	Service   string `json:"service,omitempty" binding:"required"`
	Status    string `json:"status,omitempty" binding:"required"`
	Timestamp string `json:"last_update,omitempty"`
}

// Events serves as our "status database"
var Events = make(map[string]Event)

func main() {
	router := gin.Default()

	router.Static("/public", "./public")
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/html/index.html")
	})

	router.GET("/events", GetEvents)
	router.GET("/events/:id", GetEvent)
	router.POST("/events", CreateEvent)
	router.PUT("/events/:id", UpdateEvent)

	router.Run()
}
