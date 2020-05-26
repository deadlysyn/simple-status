package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Event describes a "status event"
type Event struct {
	Service   string `json:"service,omitempty" binding:"required"`
	Status    string `json:"status,omitempty" binding:"required"`
	Timestamp string `json:"last_update,omitempty"`
}

// Events serves as our "status database"
var Events = make(map[string]Event)

// GET /events
func getEvents(c *gin.Context) {
	var events []string

	for k := range Events {
		events = append(events, k)
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

// GET /events/:id
func getEvent(c *gin.Context) {
	id := c.Param("id")

	if _, ok := Events[id]; ok {
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
	e := new(Event)

	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide service and status",
		})
		return
	}

	for k := range Events {
		if Events[k].Service == e.Service {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "service exists, use update endpoint",
			})
			return
		}
	}

	e.Timestamp = time.Now().UTC().String()
	id := strings.Replace(uuid.New().String(), "-", "", -1)
	Events[id] = *e

	c.JSON(http.StatusOK, gin.H{
		"message": "event created",
		"id":      id,
		"service": e.Service,
		"status":  e.Status,
	})
}

// PUT /events/:id
func updateEvent(c *gin.Context) {
	id := c.Param("id")

	if _, ok := Events[id]; ok {
		e := new(Event)

		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "must provide service and status",
			})
			return
		}

		e.Timestamp = time.Now().UTC().String()
		Events[id] = *e

		c.JSON(http.StatusOK, gin.H{
			"message": "event updated",
			"event":   Events[id],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "event not found",
		})
	}
}

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)

	router.Run()
}
