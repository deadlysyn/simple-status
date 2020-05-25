package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type event struct {
	Timestamp string `json:"last_update,omitempty"`
	// Service   string `json:"service,omitempty"`
	Status string `json:"status,omitempty"`
}

// Events serves as our "status database" and
// maps service names to events
var Events = make(map[string]event)

// GET /events
func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events": Events,
	})
}

// GET /events/:id
func getEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

// POST /events
func createEvent(c *gin.Context) {
	service := c.PostForm("service")
	status := c.PostForm("status")

	if service != "" {
		e := new(event)
		e.Status = string(status)
		t := time.Now().UTC()
		e.Timestamp = t.String()
		Events[service] = *e
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "event created",
		"service": service,
		"status":  status,
	})
}

// PUT /events/:id
func updateEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)

	router.Run()
}
