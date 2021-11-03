package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint       `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Categories  []Category `gorm:"many2many:product_categories;" json:"categories"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Category struct {
	ID        uint     `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Type      string   `json:"type"`
	Images    []string `gorm:"type:text[]" json:"images"`
	Price     int      `json:"price"`
	Inventory int      `json:"inventory"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
