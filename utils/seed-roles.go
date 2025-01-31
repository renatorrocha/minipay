package utils

import (
	"github.com/renatorrocha/minipay/pkg/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	// Check if roles already exist
	var count int64
	if err := db.Model(&models.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	defaultRoles := []models.Role{
		{Name: models.RoleRegular},
		{Name: models.RoleMerchant},
	}

	for _, role := range defaultRoles {
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
