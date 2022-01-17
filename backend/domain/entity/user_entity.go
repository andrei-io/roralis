package entity

import (
	"time"
)

// nolint: govet
type User struct {
	ID        uint64    `json:"ID"`
	Email     string    `json:"Email" binding:"required"`
	Name      string    `json:"Name" binding:"required"`
	Password  string    `json:"Password" binding:"required"`
	CreatedAt time.Time `json:"CreatedAt" binding:"required"`
	// DeletedAt is a pointer because it can be null
	DeletedAt *time.Time `json:"DeletedAt"`
}

// Disabled govet here.It reports that the struct can be realigned to save some bytes, but the order is ok as it is
