package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint       `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	UserID    uint       `gorm:"primaryKey"`
	CartItems []CartItem `gorm:"many2many:cart_cartItems;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CartItem struct {
	ID         uint     `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	CartID     uint     `gorm:"primaryKey"`
	ProductID  uint     `json:"product_id"`
	Product    Product  `json:"product"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
	Quantity   int      `json:"quantity"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
