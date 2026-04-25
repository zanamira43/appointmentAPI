package repository

import (
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/response"
	"gorm.io/gorm"
)

type GormDashboardRepository struct {
	DB *gorm.DB
}

func NewGormDashboardRepository(db *gorm.DB) *GormDashboardRepository {
	return &GormDashboardRepository{DB: db}
}

func (r GormDashboardRepository) DashboardOutCome() (*response.DashboardResponse, error) {

	// patient section
	var totalCountPatinet int64
	err := r.DB.Model(&models.Patient{}).Count(&totalCountPatinet).Error
	if err != nil {
		return nil, err
	}

	// session section
	var totalCountSession int64
	serr := r.DB.Model(&models.Session{}).Count(&totalCountSession).Error
	if serr != nil {
		return nil, serr
	}

	// user section
	var totalCountUser int64
	userError := r.DB.Model(&models.User{}).Count(&totalCountUser).Error
	if userError != nil {
		return nil, userError
	}

	res := &response.DashboardResponse{
		TotalPatients: totalCountPatinet,
		TotalSessions: totalCountSession,
		TotalUsers:    totalCountUser,
	}

	return res, nil
}
