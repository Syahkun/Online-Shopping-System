package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Order     []Order        `gorm:"foreignKey:DeliveryID"`
}
