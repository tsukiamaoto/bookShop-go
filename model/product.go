package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
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
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Types     pq.StringArray `gorm:"type:text[]" json:"types"`
	TypeID    *int
	Type      Type           `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"type"`
	Images    pq.StringArray `gorm:"type:text[]" json:"images"`
	Price     int            `json:"price"`
	Inventory int            `json:"inventory"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Type struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	ParentID  *int
	Parent    *Type `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent"`
	Level     int   `json:"level"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
