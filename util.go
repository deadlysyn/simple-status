package main

import (
	"strings"
	"time"
)

// TimeString returns current UTC time string
func TimeString() string {
	return time.Now().UTC().String()
}

// ValidStatus verifies given status string is valid
func ValidStatus(status string) bool {
	s := strings.ToLower(status)
	if s != "up" && s != "down" && s != "degraded" && s != "maintenance" {
		return false
	}
	return true
}
