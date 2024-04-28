package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;notn null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
}
