package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string `gorm:"size:255"`
	Username string
}
