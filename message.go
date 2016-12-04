package main

import (
	"time"
)

//message represents a single message
type message struct {
	Name    string
	Email   string
	Message string
	When    time.Time
}
