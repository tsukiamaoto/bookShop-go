package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	UserID    uint
	User      User
	CartItems []CartItem `gorm:"many2many:cart_cartItems;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type CartItem struct {
	ID        uint `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	ProductID uint
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
