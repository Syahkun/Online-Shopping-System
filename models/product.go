package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint
	UserID     uint
	CategoryID uint
	Name       string
	Price      int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User       User           `gorm:"foreignKey:UserID"`
	Category   Category       `gorm:"foreignKey:CategoryID"`
	// Order       []Order        `gorm:"foreignKey:CustomerID"`
}
