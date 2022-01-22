package entity

import (
	"time"
)

// nolint: govet
type Post struct {
	ID          uint64     `json:"ID"`
	UserID      uint64     `json:"UserID" binding:"required"`
	Latitude    float32    `json:"Latitude"`
	Longitude   float32    `json:"Longitude"`
	Title       string     `json:"Title" binding:"required"`
	Description string     `json:"Description"`
	Address     string     `json:"Address"`
	Expiry      *time.Time `json:"Expiry"`
	Photo       string     `json:"Photo"`
	RegionID    uint64     `json:"RegionID" binding:"required"`
	CategoryID  uint64     `json:"CategoryID" binding:"required"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	DeletedAt   *time.Time `json:"DeletedAt"` // DeletedAt is a pointer because it can be null
}

// Disabled govet here. It reports that the struct can be realigned to save some bytes, but the order is more logical now, if memory becomes an issue will reorder
