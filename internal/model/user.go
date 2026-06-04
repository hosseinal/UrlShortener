package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"unique"`
	Password  string
}
