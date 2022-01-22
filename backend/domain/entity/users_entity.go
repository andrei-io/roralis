package entity

import (
	"time"
)

// nolint: govet
type User struct {
	ID        uint64    `json:"ID"`
	Email     string    `json:"Email" binding:"required"`
	Password  string    `json:"Password" binding:"required"`
	Role      uint8     `json:"Role" binding:"required"`
	Name      string    `json:"Name" binding:"required"`
	Phone     string    `json:"Phone"`
	Profile   string    `json:"Profile"`
	CreatedAt time.Time `json:"CreatedAt"`
	// DeletedAt is a pointer because it can be null
	DeletedAt *time.Time `json:"DeletedAt"`
}

// Disabled govet here. It reports that the struct can be realigned to save some bytes, but the order is more logical now, if memory becomes an issue will reorder
