package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCashier struct {
	UserID    string `gorm:"size:50;not null;uniqueIndex;primary_key"`
	Username  string `gorm:"size:50;not null"`
	Password  string `gorm:"size:50;not null"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
