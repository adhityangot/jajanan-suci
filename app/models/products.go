package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID        string `gorm:"size:50;not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:255;not null"`
	Price     int    `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
