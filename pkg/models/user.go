package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string  `json:"name" gorm:"not null"`
	Email    string  `json:"email" gorm:"not null;unique"`
	Password string  `json:"password" gorm:"not null"`
	Document string  `json:"document" gorm:"not null;unique"`
	Balance  float64 `json:"balance" gorm:"not null"`
	RoleID   uint    `json:"role_id" gorm:"not null"`
	Role     Role    `json:"role" gorm:"foreignKey:RoleID"`
}
