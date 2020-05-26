package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetEvents provides list of event IDs
// GET /events
func GetEvents(c *gin.Context) {
	var events []string

	for k := range Events {
		events = append(events, k)
	}

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

// GetEvent GET /events/:id
func GetEvent(c *gin.Context) {
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

// CreateEvent creates a new status event
// POST /events
// {service: foo, status: bar}
func CreateEvent(c *gin.Context) {
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

// UpdateEvent updates an existing status event
// PUT /events/:id
// {service: foo, status: bar}
func UpdateEvent(c *gin.Context) {
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
