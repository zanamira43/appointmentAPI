package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormSessionRepository struct {
	DB *gorm.DB
}

func NewGormSessionRepository(db *gorm.DB) *GormSessionRepository {
	return &GormSessionRepository{DB: db}
}

// insert new offer data into sql database
func (r *GormSessionRepository) CreateSession(session *dto.Session) error {
	return r.DB.Create(&session).Error
}

// get all sessions data from sql database
func (r *GormSessionRepository) GetAllSessions(page, limit int) ([]models.Session, int64, error) {
	var sessions []models.Session

	var total int64

	// get total number of sessions
	err := r.DB.Model(&sessions).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err = r.DB.Offset(offset).Limit(limit).Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return sessions, total, nil
}

// get offer data by id from sql database
func (r *GormSessionRepository) GetSessionByID(id uint) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("id = ?", id).First(&session).Preload("Patient").Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// update offer data by id from sql database
func (r *GormSessionRepository) UpdateSession(id uint, sessionDto *dto.Session) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}

	if sessionDto.PatientID != 0 {
		session.PatientID = sessionDto.PatientID
	}
	if session.Duration != 0 {
		session.Duration = sessionDto.Duration
	}
	if sessionDto.Status != "" {
		session.Status = sessionDto.Status
	}
	if sessionDto.Notes != "" {
		session.Notes = sessionDto.Notes
	}
	if sessionDto.SessionDate != "" {
		session.SessionDate = sessionDto.SessionDate
	}

	err = r.DB.Save(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// delete offer data by id from sql database
func (r *GormSessionRepository) DeleteSession(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.Session{}).Error
}
