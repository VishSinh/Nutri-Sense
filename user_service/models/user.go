package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"uniqueIndex;notn null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"uniqueIndex;not null"`
}

type UserDetails struct {
	UserID uuid.UUID `gorm:"primaryKey;foreignKey:ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name   string    `gorm:"not null"`
	Age    int       `gorm:"not null"`
	Weight float32   `gorm:"not null"`
	Height int       `gorm:"null"`
}
