package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID              uint       `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Categories      []Category `gorm:"many2many:product_categories;" json:"categories"`
	Editor          string     `json:"editor"`
	Publisher       string     `json:"publisher"`
	PublicationDate string     `json:"publicaionDate"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Category struct {
	ID        uint           `gorm:"primaryKey;uniqueIndex;autoIncrement"`
	Types     pq.StringArray `gorm:"type:text[]"`
	Images    pq.StringArray `gorm:"type:text[]" json:"images"`
	Price     int            `json:"price"`
	Inventory int            `json:"inventory"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
