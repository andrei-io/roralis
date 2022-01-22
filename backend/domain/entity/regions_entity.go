package entity

import (
	"time"
)

// nolint: govet
type Region struct {
	ID        uint64    `json:"ID"`
	Text      string    `json:"Text" binding:"required"`
	CreatedAt time.Time `json:"CreatedAt"`
	// DeletedAt is a pointer because it can be null
	DeletedAt *time.Time `json:"DeletedAt"`
}

// Disabled govet here. It reports that the struct can be realigned to save some bytes, but the order is more logical now, if memory becomes an issue will reorder
