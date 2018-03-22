package main

import (
	"time"
)

// message represents a single message
type message struct {
	Name      string
	Message   string
	AvatarURL string
	When      time.Time
}
