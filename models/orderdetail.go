package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Amount    int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Order     Order          `gorm:"foreignKey:OrderID"`
	Product   Product        `gorm:"foreignKey:ProductID"`
}
