package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint        `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	OderItems []OrderItem `gorm:"many2many:order_orderItems;"`
	UserID    uint
	User      User
	Total     int `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	ProductID uint
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
