package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint        `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	OrderItems []OrderItem `gorm:"many2many:order_orderItems;"`
	UserID     uint        `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type OrderItem struct {
	ID         uint     `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	OrderID    uint     `gorm:"primaryKey"`
	ProductID  uint     `json:"product_id"`
	Product    Product  `json:"product"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
	Quantity   int      `json:"quantity"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
