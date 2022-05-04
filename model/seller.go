package model

import (
	"time"

	"gorm.io/gorm"
)

type Seller struct {
	ID        uint `gorm:"primaryKey;uniqueIndex;autoIncrement" json:"id"`
	UserID    uint
	Products  []*Product `gorm:"many2many:seller_products"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
