package models

import "gorm.io/gorm"

type UserDetail struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:32"`
	PasswordHash string `json:"password_hash"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}
