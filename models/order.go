package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint
	UserID     uint
	DeliveryID uint
	TotalPrice int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User       User           `gorm:"foreignKey:UserID"`
	Delivery   Delivery       `gorm:"foreignKey:DeliveryID"`
}
