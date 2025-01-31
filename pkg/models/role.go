package models

import "gorm.io/gorm"

type RoleType string

const (
	RoleRegular  RoleType = "Regular"
	RoleMerchant RoleType = "Merchant"
)

type Role struct {
	gorm.Model
	Name  RoleType `json:"name" gorm:"not null;type:varchar(20)"`
	Users []User   `json:"users" gorm:"foreignKey:RoleID"`
}
