package utils

import (
	"github.com/renatorrocha/minipay/pkg/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
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
