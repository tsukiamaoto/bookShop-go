package model

import (
	"time"

	"gorm.io/gorm"
)

// user model
type User struct {
	ID        uint   `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
