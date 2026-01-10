package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormPersonInfoRepository struct {
	DB *gorm.DB
}

func NewGormPersonInfoRepository(db *gorm.DB) *GormPersonInfoRepository {
	return &GormPersonInfoRepository{DB: db}
}

// create new person info
func (r *GormPersonInfoRepository) CreatePersonInfo(personInfo *dto.PersonInfo) error {
	return r.DB.Create(&personInfo).Error
}

// get list of person info
func (r *GormPersonInfoRepository) GetAllPersonInfo(page, limit int, search string) ([]models.PersonInfo, int64, error) {
	var personInfo []models.PersonInfo
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.PersonInfo{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("full_name LIKE ? OR phone_number LIKE ?", searchPattern, searchPattern)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Order("id DESC").Find(&personInfo).Error
	if err != nil {
		return nil, 0, err
	}
	return personInfo, total, nil
}

// get single person info
func (r *GormPersonInfoRepository) GetPersonInfoByID(id uint) (*models.PersonInfo, error) {
	var personInfo models.PersonInfo
	err := r.DB.Where("id = ?", id).First(&personInfo).Error
	if err != nil {
		return nil, err
	}
	return &personInfo, nil
}

// get single person info by phone. number
func (r *GormPersonInfoRepository) GetPersonInforByPhone(phone_number string) (*models.PersonInfo, error) {
	var personInfo models.PersonInfo
	err := r.DB.Where("phone_number = ?", phone_number).First(&personInfo).Error
	if err != nil {
		return nil, err
	}
	return &personInfo, nil
}

// update single person info
func (r *GormPersonInfoRepository) UpdatePersonInfo(id uint, updatePersonInfo *dto.PersonInfo) (*models.PersonInfo, error) {
	personInfo, err := r.GetPersonInfoByID(id)
	if err != nil {
		return nil, err
	}

	if updatePersonInfo.FullName != "" {
		personInfo.FullName = updatePersonInfo.FullName
	}
	if updatePersonInfo.PhoneNumber != "" {
		personInfo.PhoneNumber = updatePersonInfo.PhoneNumber
	}

	err = r.DB.Save(&personInfo).Error
	if err != nil {
		return nil, err
	}
	return personInfo, nil
}

// delete single person info
func (r *GormPersonInfoRepository) DeletePersonInfo(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&models.PersonInfo{}).Error
}
