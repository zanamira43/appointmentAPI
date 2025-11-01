package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormSettingsRepository struct {
	DB *gorm.DB
}

func NewGormSettingsRepository(db *gorm.DB) GormSettingsRepository {
	return GormSettingsRepository{DB: db}
}

// get single  setting
func (r *GormSettingsRepository) GetSetting() (*models.Settings, error) {
	var setting models.Settings

	err := r.DB.Find(&setting).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

// update setting
func (r *GormSettingsRepository) UpdateSetting(dto *dto.Settings) (*models.Settings, error) {
	var setting models.Settings

	if dto.ID != 0 {
		setting.ID = dto.ID
	}

	if dto.SystemName != setting.SystemName {
		setting.SystemName = dto.SystemName
	}

	if dto.PhoneNumber != setting.PhoneNumber {
		setting.PhoneNumber = dto.PhoneNumber
	}

	if dto.Address != setting.Address {
		setting.Address = dto.Address
	}

	if dto.BillPrefix != setting.BillPrefix {
		setting.BillPrefix = dto.BillPrefix
	}

	if err := r.DB.Save(&setting).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
