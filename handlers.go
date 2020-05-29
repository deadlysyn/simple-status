package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvents provides list of all events
// GET /events
func GetEvents(c *gin.Context) {
	var l []Event
	for k := range Events {
		l = append(l, Event{
			Service:     k,
			Status:      Events[k].Status,
			Description: Events[k].Description,
			Updated:     Events[k].Updated,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"events": l,
	})
}

// GetEvent provides event detail for specified service
// GET /events/:service
func GetEvent(c *gin.Context) {
	s := c.Param("service")
	if _, ok := Events[s]; ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "service found",
			s:         Events[s],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "service not found",
		})
	}
}

// CreateEvent creates a new status event
//
// POST /events
// {
// service: name,
// status: up|down|degraded|maintenance,
// description: text
// }
func CreateEvent(c *gin.Context) {
	var e Event
	if err := c.ShouldBindJSON(&e); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide service, status and description",
		})
		return
	}

	if _, ok := Events[e.Service]; ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "service exists, use PUT",
		})
		return
	}

	if !ValidStatus(e.Status) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid status",
		})
		return
	}

	Events[e.Service] = EventData{
		Status:      e.Status,
		Description: e.Description,
		Updated:     TimeString(),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "service created",
		e.Service: Events[e.Service],
	})
}

// UpdateEvent updates an existing event
// PUT /events/:service
// {
// status: up|down|degraded|maintenance,
// description: text
//}
func UpdateEvent(c *gin.Context) {
	s := c.Param("service")
	if _, ok := Events[s]; ok {
		var e EventData
		if err := c.ShouldBindJSON(&e); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "must provide status and description",
			})
			return
		}

		if !ValidStatus(e.Status) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid status",
			})
			return
		}

		e.Updated = TimeString()
		Events[s] = e

		c.JSON(http.StatusOK, gin.H{
			"message": "status updated",
			s:         Events[s],
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "service not found",
		})
	}
}
