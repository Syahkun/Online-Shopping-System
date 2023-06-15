package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name        string
	Username    string `gorm:"type:varchar(300);not null;unique"`
	Password    string `gorm:"type:varchar(300);not null"`
	Email       *string
	Address     string
	PhoneNumber string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Order       []Order        `gorm:"foreignKey:CustomerID"`
}
