package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type event struct {
	Timestamp string `json:"last_update,omitempty"`
	Service   string `json:"service,omitempty"`
	Status    string `json:"status,omitempty"`
}

// Events serves as our "status database"
var Events = make(map[string]event)

// GET /events
func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events": Events,
	})
}

// GET /events/:id
func getEvent(c *gin.Context) {
	id := c.Param("id")

	_, ok := Events[id]
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "event found",
			"event":   Events[id],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
	}
}

// POST /events
func createEvent(c *gin.Context) {
	service := c.PostForm("service")
	status := c.PostForm("status")

	if service == "" || status == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide service and status",
		})
		return
	}

	for k := range Events {
		if Events[k].Service == service {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "service exists, use update endpoint",
			})
			return
		}
	}

	e := new(event)

	t := time.Now().UTC()
	e.Timestamp = t.String()

	e.Service = string(service)
	e.Status = string(status)

	id := strings.Replace(uuid.New().String(), "-", "", -1)
	Events[id] = *e

	c.JSON(http.StatusOK, gin.H{
		"message": "event created",
		"id":      id,
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
