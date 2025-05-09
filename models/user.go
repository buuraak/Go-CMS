package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey"`
	Username          string `gorm:"unique;not null"`
	Email             string `gorm:"unique;not null"`
	FirstName         string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	Password          string `gorm:"not null" json:"-"`
	Role              string `gorm:"type:enum('admin','editor','customer');default:'customer';not null"`
	IsVerified        bool   `gorm:"default:false;not null"`
	VerificationToken string
}
